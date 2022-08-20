package deepcopy

func ProduceDeepCopyInterface(config *Config, origin interface{}) interface{} {
	if config == nil {
		config = NewDefaultConfig()
	}

	return produceInterface(config, origin)
}
