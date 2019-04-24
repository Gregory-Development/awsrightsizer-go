package app

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type AWSInstances struct {
	EC2Instances ec2Instances
	RDSInstances rdsInstances
}

type ec2Instances struct {
	GeneralPurpose ec2GeneralPurposeInstances `yaml:"GeneralPurpose"`
	ComputeOptimized ec2ComputeOptimizedInstances `yaml:"ComputeOptimized"`
	MemoryOptimized ec2MemoryOptimizedInstances `yaml:"MemoryOptimized"`
	AcceleratedComputing ec2AcceleratedComputingInstances `yaml:"AcceleratedComputing"`
	FPGAAcceleratedComputing ec2FPGAAcceleratedComputingInstances `yaml:"FPGAAcceleratedComputing"`
	StorageOptimized ec2StorageOptimizedInstances `yaml:"StorageOptimized"`
}

type rdsInstances struct {
	GeneralPurpose rdsGeneralPurposeInstances `yaml:"GeneralPurpose"`
	MemoryOptimized rdsMemoryOptimizedInstances `yaml:"MemoryOptimized"`
}

type GeneralInstance struct {
	Name string `yaml:"name"`
	Vcpu int `yaml:"vcpu"`
	CpuCreditPerHour int `yaml:"cpu_credit_per_hour"`
	Memory float32 `yaml:"memory"`
	Storage struct {
		Type string `yaml:"type"`
		Qty int `yaml:"qty"`
		Size int `yaml:"size"`
		MaxDedicatedEBSBandwidth int `yaml:"max_dedicated_ebs_bandwidth"`
	} `yaml:"storage"`
	NetMax int `yaml:"net_max"`
}

type AcceleratedInstance struct {
	Name string `yaml:"name"`
	Gpu struct {
		Count int `yaml:"count"`
		Type string `yaml:"type"`
		Memory int `yaml:"memory"`
	} `yaml:"gpu"`
	Vcpu int `yaml:"vcpu"`
	CpuCreditPerHour int `yaml:"cpu_credit_per_hour"`
	Memory float32 `yaml:"memory"`
	Storage struct {
		Type string `yaml:"type"`
		Qty int `yaml:"qty"`
		Size int `yaml:"size"`
		MaxDedicatedEBSBandwidth int `yaml:"max_dedicated_ebs_bandwidth"`
	} `yaml:"storage"`
	NetMax int `yaml:"net_max"`
}

type ec2GeneralPurposeInstances struct {
	A1 []GeneralInstance `yaml:"a1"`
	T3 []GeneralInstance `yaml:"t3"`
	T3a []GeneralInstance `yaml:"t3a"`
	T2 []GeneralInstance `yaml:"t2"`
	T1 []GeneralInstance `yaml:"t1"`
	M5 []GeneralInstance `yaml:"m5"`
	M5d []GeneralInstance `yaml:"m5d"`
	M5a []GeneralInstance `yaml:"m5a"`
	M5ad []GeneralInstance `yaml:"m5ad"`
	M4 []GeneralInstance `yaml:"m4"`
	M3 []GeneralInstance `yaml:"m3"`
	M1 []GeneralInstance `yaml:"m1"`
}

type ec2ComputeOptimizedInstances struct {
	C5 []GeneralInstance `yaml:"c5"`
	C5d []GeneralInstance `yaml:"c5d"`
	C5n []GeneralInstance `yaml:"c5n"`
	C4 []GeneralInstance `yaml:"c4"`
	C3 []GeneralInstance `yaml:"c3"`
	CC2 []GeneralInstance `yaml:"cc2"`
	C1 []GeneralInstance `yaml:"c1"`
}

type ec2MemoryOptimizedInstances struct {
	R5 []GeneralInstance `yaml:"r5"`
	R5d []GeneralInstance `yaml:"r5d"`
	R5a []GeneralInstance `yaml:"r5a"`
	R5ad []GeneralInstance `yaml:"r5ad"`
	R4 []GeneralInstance `yaml:"r4"`
	R3 []GeneralInstance `yaml:"r3"`
	CR1 []GeneralInstance `yaml:"cr1"`
	M2 []GeneralInstance `yaml:"m2"`
	X1e []GeneralInstance `yaml:"x1e"`
	X1 []GeneralInstance `yaml:"x1"`
	U6tb1 []GeneralInstance `yaml:"u-6tb1"`
	U9tb1 []GeneralInstance `yaml:"u-9tb1"`
	U12tb1 []GeneralInstance `yaml:"u-12tb1"`
	Z1d []GeneralInstance `yaml:"z1d"`
}

type ec2AcceleratedComputingInstances struct {
	P3 []AcceleratedInstance `yaml:"p3"`
	P3dn []AcceleratedInstance `yaml:"p3dn"`
	P2 []AcceleratedInstance `yaml:"p2"`
	G3 []AcceleratedInstance `yaml:"g3"`
	G3s []AcceleratedInstance `yaml:"g3s"`
	G2 []AcceleratedInstance `yaml:"g2"`
}

type ec2FPGAAcceleratedComputingInstances struct {
	F1 []struct {
		Name string `yaml:"name"`
		FPGACount int `yaml:"fpga_count"`
		Vcpu int `yaml:"vcpu"`
		CpuCreditPerHour int `yaml:"cpu_credit_per_hour"`
		Memory float32 `yaml:"memory"`
		Storage struct {
			Type string `yaml:"type"`
			Qty int `yaml:"qty"`
			Size int `yaml:"size"`
			MaxDedicatedEBSBandwidth int `yaml:"max_dedicated_ebs_bandwidth"`
		} `yaml:"storage"`
		NetMax int `yaml:"net_max"`
	} `yaml:"f1"`
}

type ec2StorageOptimizedInstances struct {
	H1 []GeneralInstance `yaml:"h1"`
	I3 []GeneralInstance `yaml:"i3"`
	I2 []GeneralInstance `yaml:"i2"`
	D2 []GeneralInstance `yaml:"d2"`
	HS1 []GeneralInstance `yaml:"hs1"`
}

type rdsGeneralPurposeInstances struct {
	T3 []GeneralInstance `yaml:"t3"`
	T2 []GeneralInstance `yaml:"t2"`
	M5 []GeneralInstance `yaml:"m5"`
	M4 []GeneralInstance `yaml:"m4"`
	M3 []GeneralInstance `yaml:"m3"`
	M1 []GeneralInstance `yaml:"m1"`
}

type rdsMemoryOptimizedInstances struct {
	R5 []GeneralInstance `yaml:"r5"`
	R4 []GeneralInstance `yaml:"r4"`
	R3 []GeneralInstance `yaml:"r3"`
	X1e []GeneralInstance `yaml:"x1e"`
	X1 []GeneralInstance `yaml:"x1"`
	M2 []GeneralInstance `yaml:"m2"`
}

func NewInstances() *AWSInstances {
	var a AWSInstances

	a.newEc2Instances()
	a.newRdsInstances()

	return &a
}

func (a *AWSInstances) newEc2Instances() {
	f, err := ioutil.ReadFile("config/ec2-instances.yaml")
	if err != nil {
		log.Println(err)
	}

	log.Println(yaml.Unmarshal(f, &a.EC2Instances))
}

func (a *AWSInstances) newRdsInstances() {
	f, err := ioutil.ReadFile("config/rds-instances.yaml")
	if err != nil {
		log.Println(err)
	}

	log.Println(yaml.Unmarshal(f, &a.RDSInstances))
}