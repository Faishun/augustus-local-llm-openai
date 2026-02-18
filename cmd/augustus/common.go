package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/praetorian-inc/augustus/pkg/buffs"
	"github.com/praetorian-inc/augustus/pkg/detectors"
	"github.com/praetorian-inc/augustus/pkg/generators"
	"github.com/praetorian-inc/augustus/pkg/harnesses"
	"github.com/praetorian-inc/augustus/pkg/probes"
	"github.com/praetorian-inc/augustus/pkg/registry"
	"github.com/praetorian-inc/augustus/pkg/types"
)

const version = "0.1.0"

func listCapabilities() {
	fmt.Println("Registered Capabilities")
	fmt.Println("=======================")
	fmt.Println()

	fmt.Printf("Probes (%d):\n", probes.Registry.Count())
	for _, name := range probes.List() {
		fmt.Printf("  - %s\n", name)
	}
	fmt.Println()

	fmt.Printf("Generators (%d):\n", generators.Registry.Count())
	for _, name := range generators.List() {
		fmt.Printf("  - %s\n", name)
	}
	fmt.Println()

	fmt.Printf("Detectors (%d):\n", detectors.Registry.Count())
	for _, name := range detectors.List() {
		fmt.Printf("  - %s\n", name)
	}
	fmt.Println()

	fmt.Printf("Harnesses (%d):\n", harnesses.Registry.Count())
	for _, name := range harnesses.List() {
		fmt.Printf("  - %s\n", name)
	}
	fmt.Println()

	fmt.Printf("Buffs (%d):\n", buffs.Registry.Count())
	for _, name := range buffs.List() {
		fmt.Printf("  - %s\n", name)
	}
}

// listCapabilitiesJSON outputs probes as JSON: {"probeName": {"description": "...", "goal": "..."}, ...}.
// Description and goal are empty when the probe does not implement ProbeMetadata.
func listCapabilitiesJSON() {
	type meta struct {
		Description string `json:"description"`
		Goal        string `json:"goal"`
	}
	out := make(map[string]meta)
	for _, name := range probes.List() {
		m := meta{}
		probe, err := probes.Create(name, registry.Config{})
		if err == nil {
			if pm, ok := probe.(types.ProbeMetadata); ok {
				m.Description = pm.Description()
				m.Goal = pm.Goal()
			}
		}
		out[name] = m
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(out)
}
