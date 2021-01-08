package commonprofiler

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
)

type Service struct {
	cfg *ProfilerCliConfig
	cf  *os.File
}

func New(cfg *ProfilerCliConfig) *Service {
	return &Service{cfg: cfg}
}

func (s *Service) Begin() error {
	if s.cfg.CpuProfile != "" {
		return s.startCpuProfile()
	}

	return nil
}

func (s *Service) startCpuProfile() (err error) {
	s.cf, err = os.Create(s.cfg.CpuProfile)
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}
	if err := pprof.StartCPUProfile(s.cf); err != nil {
		return fmt.Errorf("start cpu profile: %w", err)
	}

	return nil
}

func (s *Service) End() error {
	var err error

	if s.cfg.CpuProfile != "" {
		err = s.stopCpuProfile()
	}
	if err != nil {
		return fmt.Errorf("stop cpu profile: %w", err)
	}

	if s.cfg.MemProfile != "" {
		err = s.writeMemProfile()
	}
	if err != nil {
		return fmt.Errorf("write mem profile: %w", err)
	}

	return nil
}

func (s *Service) stopCpuProfile() error {
	pprof.StopCPUProfile()
	if err := s.cf.Close(); err != nil {
		return fmt.Errorf("close: %w", err)
	}

	return nil
}

func (s *Service) writeMemProfile() error {
	f, err := os.Create(s.cfg.MemProfile)
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}

	runtime.GC()
	if err := pprof.WriteHeapProfile(f); err != nil {
		return fmt.Errorf("write heap profile: %w", err)
	}

	return f.Close()
}
