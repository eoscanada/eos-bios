package bios

import (
	"fmt"
	"strings"

	"github.com/eoscanada/eos-bios/bios/disco"
)

func ComputeContentsAgreement(orderedPeers []*Peer) (out *ConsensusFiles) {
	/**
		  ConsensusFiles{
		    FilesMap: {
	          "boot_sequence.yaml": ConsensusRefs{
		        "/ipfs/Qm123": ContentRefVersion{Version: 1, ref},
		        "/ipfs/Qm234": ContentRefVersion{Version: 2, ref},
		      },
	        },
	        FilesList: []string{"boot_sequence.yaml"},
		  }
	*/
	out = &ConsensusFiles{
		FilesMap: make(map[string]ConsensusRefs),
	}

	for _, peer := range orderedPeers {
		for _, contentRef := range peer.Discovery.TargetContents {

			if out.FilesMap[contentRef.Name] == nil {
				out.FilesMap[contentRef.Name] = ConsensusRefs{}
				out.FilesList = append(out.FilesList, contentRef.Name)
			}

			thisFileMap := out.FilesMap[contentRef.Name]

			if thisFileMap[contentRef.Ref] == 0 {
				thisFileMap[contentRef.Ref] = len(thisFileMap) + 1
			}
		}
	}

	return
}

type ConsensusFiles struct {
	FilesList []string // []string{"boot_sequence.yaml", "snapshot.csv"}
	FilesMap  map[string]ConsensusRefs
}

// FIXME: the last bit here could simply be a version number.. it'd be fine..
type ConsensusRefs map[string]int

type ContentRefVersion struct {
	disco.ContentRef
	Version int
}

func ComputePeerContentsColumn(content *ConsensusFiles, orderedPeers []*Peer) (out []string) {
	/*
		[
			"1111211",
			"2111111",
			"2111211",
			"1111111",
		]
	*/

	for _, peer := range orderedPeers {
		var currentColumns []string
		for _, fileName := range content.FilesList {
			col := "."
			for _, ref := range peer.Discovery.TargetContents {
				if ref.Name == fileName {
					version := content.FilesMap[fileName][ref.Ref]
					if version > 9 {
						col = "X"
					} else {
						col = fmt.Sprintf("%d", version)
					}
					break
				}
			}
			currentColumns = append(currentColumns, col)
		}
		out = append(out, strings.Join(currentColumns, ""))
	}

	return
}
