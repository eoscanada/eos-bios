package bios

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	eos "github.com/eoscanada/eos-go"
)

func Serve(net *Network) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("DEV") != "" {
			cnt, err := ioutil.ReadFile("/home/abourget/go/src/github.com/eoscanada/eos-bios/visualization/d3viz.html")
			if err != nil {
				fmt.Println("error reading d3viz.html:", err)
				return
			}

			w.Write(cnt)
		} else {
			_, _ = w.Write([]byte(vizTemplate))
		}
	})
	http.HandleFunc("/data.json", func(w http.ResponseWriter, r *http.Request) {
		pov := r.FormValue("pov")
		fmt.Println("Serving request for /data.json from point of view of account", pov)

		out := struct {
			Nodes []*vizNode `json:"nodes"`
			Links []*vizEdge `json:"links"`
		}{}

		network := net.NetworkThatIncludes(eos.AccountName(pov))
		if network != nil {
			for _, node := range network.Nodes() {
				out.Nodes = append(out.Nodes, &vizNode{
					ID:          string(node.(*Peer).Discovery.SeedNetworkAccountName),
					TotalWeight: node.(*Peer).TotalWeight,
				})
			}

			for _, edge := range network.WeightedEdges() {
				out.Links = append(out.Links, &vizEdge{
					Source: string(edge.From().(*Peer).Discovery.SeedNetworkAccountName),
					Target: string(edge.To().(*Peer).Discovery.SeedNetworkAccountName),
					Weight: edge.Weight(),
				})
			}
		}

		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")
		enc.Encode(out)
	})

	fmt.Println("Serving visualization on http://127.0.0.1:10101 ...")
	if err := http.ListenAndServe("0.0.0.0:10101", nil); err != nil {
		fmt.Println("ERROR listening on network visualition endpoint:", err)
	}
}

type vizNode struct {
	ID          string `json:"id"`
	TotalWeight int    `json:"total_weight"`
}
type vizEdge struct {
	Source string  `json:"source"`
	Target string  `json:"target"`
	Weight float64 `json:"weight"`
}

var vizTemplate = `
<!DOCTYPE html>
<html><head>
<meta charset="utf-8">
<style>

.links line {
  stroke: #999;
  stroke-opacity: 0.6;
}

.nodes circle {
  stroke: #fff;
  stroke-width: 1.5px;
}

</style>
</head>
<body>

<h2>Visualization of the network topology</h2>
<p>From the point of view of account: <strong id="pov"></strong>.</p>
<p>Use query param '?pov=account' to change point of view.</p>

<svg width="960" height="600"></svg>
<script src="https://d3js.org/d3.v4.min.js"></script>
<script src="https://d3js.org/d3-scale-chromatic.v1.min.js"></script>
<script>

var i = 0;
var telem;
var search_values=location.search.replace('\?','').split('&');
var query={}
for(i=0;i<search_values.length;i++){
    telem=search_values[i].split('=');
    query[telem[0]]=telem[1];
}

var svg = d3.select("svg"),
    width = +svg.attr("width"),
    height = +svg.attr("height");

var color = d3.scaleOrdinal(d3.schemeSet1);

var simulation = d3.forceSimulation()
    .force("link", d3.forceLink().id(function(d) { return d.id; }))
    .force("charge", d3.forceManyBody().strength(-250))
    .force("center", d3.forceCenter(width / 2, height / 2))
    .force('collision', d3.forceCollide().radius(function(d) {
      return d.weight;
    }));

var pov = query.pov ? query.pov : "eoscanadacom";
document.querySelector("#pov").innerHTML = pov;

d3.json("/data.json?pov=" + pov, function(error, graph) {
  if (error) throw error;

  var link = svg.append("g")
      .attr("class", "links")
    .selectAll("line")
    .data(graph.links)
    .enter().append("line")
      .attr("stroke-width", 1.2); // function(d) { return Math.sqrt(d.weight); });

  var node = svg.append("g")
      .attr("class", "nodes")
    .selectAll("circle")
    .data(graph.nodes)
    .enter().append("circle")
      .attr("r", function(d) { return Math.sqrt(d.total_weight) + 2; })
      .attr("fill", function(d, i) {
          var count = 0;
          for (i = 0; i < d.id.length; i++) {
              count += d.id.charCodeAt(i)
          }
          return color(count % graph.nodes.length);
      })
      .call(d3.drag()
          .on("start", dragstarted)
          .on("drag", dragged)
          .on("end", dragended));

  node.append("title")
      .text(function(d) { return d.id + " (" + d.total_weight + ")"; });

  simulation
      .nodes(graph.nodes)
      .on("tick", ticked);

  simulation.force("link")
      .links(graph.links);

  function ticked() {
    link
        .attr("x1", function(d) { return d.source.x; })
        .attr("y1", function(d) { return d.source.y; })
        .attr("x2", function(d) { return d.target.x; })
        .attr("y2", function(d) { return d.target.y; });

    node
        .attr("cx", function(d) { return d.x; })
        .attr("cy", function(d) { return d.y; });
  }
});

function dragstarted(d) {
  if (!d3.event.active) simulation.alphaTarget(0.3).restart();
  d.fx = d.x;
  d.fy = d.y;
}

function dragged(d) {
  d.fx = d3.event.x;
  d.fy = d3.event.y;
}

function dragended(d) {
  if (!d3.event.active) simulation.alphaTarget(0);
  d.fx = null;
  d.fy = null;
}

</script>


</body>
</html>
`
