package bios

// Discovery is data published autonomously by Block Producer
// candidates. LaunchData will contain the topology and build up the
// graph.  Otherwise, this data is simply metadata about a BP.
//type Discovery struct {
//	// Testnet is true if this discovery file represents a testing
//	// network.
//	Testnet bool `json:"testnet"`
//	// Mainnet is true if this discovery file represents the main net
//	// (or a production network). One of Testnet and Mainnet must be
//	// `true`, and are mutually exclusive.
//	Mainnet bool `json:"mainnet"`
//
//	EOSIOAccountName      string        `json:"eosio_account_name"`
//	EOSIOP2P              string        `json:"eosio_p2p"`
//	EOSIOHTTP             string        `json:"eosio_http"`
//	EOSIOHTTPS            string        `json:"eosio_https"`
//	EOSIOABPSigningKey    ecc.PublicKey `json:"eosio_appointed_block_producer_signing_key"`
//	EOSIOInitialAuthority struct {
//		Owner    eos.Authority `json:"owner"`
//		Active   eos.Authority `json:"active"`
//		Recovery eos.Authority `json:"recovery"`
//	} `json:"eosio_initial_authority"`
//
//	Website             string `json:"website"`
//	IntroductionPostURL string `json:"introduction_post_url"`
//	SocialFacebook      string `json:"social_facebook"`
//	SocialTwitter       string `json:"social_twitter"`
//	SocialYouTube       string `json:"social_youtube"`
//	SocialTelegram      string `json:"social_telegram"`
//	SocialWeChat        string `json:"social_wechat"`
//	SocialSlack         string `json:"social_slack"`
//	SocialSteemIt       string `json:"social_steemit"`
//	SocialSteem         string `json:"social_steem"`
//	SocialGitHub        string `json:"social_github"`
//	SocialKeybase       string `json:"social_keybase"`
//
//	OrganizationName    string `json:"organization_name"`
//	OrganizationTagline string `json:"organization_tagline"`
//
//	LaunchData LaunchData `json:"launch_data"`
//
//	ClonedFrom string `json:"-"`
//}
