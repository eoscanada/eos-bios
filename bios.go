package main

type BIOS struct {
	LaunchData        *LaunchData
	Config            *Config
	ShuffledProducers []*ProducerData
}

func NewBIOS(launchData *LaunchData, config *Config) *BIOS {
	b := &BIOS{
		LaunchData: launchData,
		Config:     config,
	}
	return b
}

func (b *BIOS) ShuffleProducers(btcMerkleRoot []byte) error {
	// we'll shuffle later :)
	if b.Config.NoShuffle {
		b.ShuffledProducers = b.LaunchData.Producers
	} else {
		// FIXME: put an algorithm here..
		b.ShuffledProducers = b.LaunchData.Producers
	}
	return nil
}

func (b *BIOS) IsBootNode(account string) bool {
	return b.ShuffledProducers[0].EOSIOAccountName == account
}

func (b *BIOS) AmIBootNode() bool {
	return b.IsBootNode(b.Config.Producer.MyAccount)
}

func (b *BIOS) IsAppointedBlockProducer(account string) bool {
	for i := 1; i < 22 && len(b.ShuffledProducers) > i; i++ {
		if b.ShuffledProducers[i].EOSIOAccountName == account {
			return true
		}
	}
	return false
}

func (b *BIOS) AmIAppointedBlockProducer() bool {
	return b.IsAppointedBlockProducer(b.Config.Producer.MyAccount)
}
