package app

type EC2InstanceTypes struct {
	GeneralPurpose struct {
		A1 []map[string]map[string]string
		T1 []map[string]map[string]string
		T2 []map[string]map[string]string
		T3 []map[string]map[string]string
		T3a []map[string]map[string]string
		M4 []map[string]map[string]string
		M5 []map[string]map[string]string
		M5a []map[string]map[string]string
	}
	ComputeOptimized struct {
		C4 []map[string]map[string]string
		C5 []map[string]map[string]string
		C5n []map[string]map[string]string
	}
	MemoryOptimized struct {
		R4 []map[string]map[string]string
		R5 []map[string]map[string]string
		R5a []map[string]map[string]string
		X1 []map[string]map[string]string
		X1e []map[string]map[string]string
		HighMemory []map[string]map[string]string
		Z1d []map[string]map[string]string
	}
	AcceleratedComputing struct {
		P2 []map[string]map[string]string
		P3 []map[string]map[string]string
		G3 []map[string]map[string]string
		F1 []map[string]map[string]string
	}
	StorageOptimized struct {
		H1 []map[string]map[string]string
		I3 []map[string]map[string]string
		D2 []map[string]map[string]string
	}
}
