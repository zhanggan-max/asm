package main

import (
	"asm/internal/commands"
	"fmt"
	"os"
)

func usage() {
	fmt.Println("asm - Agent Skill Manager")
	fmt.Println("\nUsage:")
	fmt.Println("  asm <command> [arguments]")
	fmt.Println("\nCommands:")
	fmt.Println("  init          Initialize a new skill project")
	fmt.Println("  install [url] Install dependencies or a specific package")
	fmt.Println("  remove <name> Remove a specific package")
	fmt.Println("  list          List installed packages")
	fmt.Println("  help          Show this help message")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	command := os.Args[1]

	// Simple routing
	switch command {
	case "init":
		// Could handle flags for init here if needed
		if err := commands.RunInit(""); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	case "install":
		// Handle arguments for install
		// asm install -> installAll
		// asm install <url> -> installPackage
		pkgURL := ""
		if len(os.Args) > 2 {
			pkgURL = os.Args[2]
		}
		if err := commands.RunInstall(pkgURL); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	case "remove":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "Error: package name required\n")
			os.Exit(1)
		}
		pkgName := os.Args[2]
		if err := commands.RunRemove(pkgName); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	case "list":
		if err := commands.RunList(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	case "help":
		usage()

	default:
		fmt.Printf("Unknown command: %s\n", command)
		usage()
		os.Exit(1)
	}
}
