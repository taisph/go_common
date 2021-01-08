package commonprofiler

type ProfilerCliConfig struct {
	CpuProfile string `name:"cpuprofile" help:"Write cpu profile to FILE." placeholder:"FILE"`
	MemProfile string `name:"memprofile" help:"Write memory profile to FILE." placeholder:"FILE"`
}

type ProfilerCli struct {
	ProfilerCliConfig `embed`
}
