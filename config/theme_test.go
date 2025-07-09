package config

import (
	"testing"
)

// Official DaisyUI themes from https://daisyui.com/docs/themes/
var validDaisyUIThemes = []string{
	"light", "dark", "cupcake", "bumblebee", "emerald", "corporate", "synthwave",
	"retro", "cyberpunk", "valentine", "halloween", "garden", "forest", "aqua",
	"lofi", "pastel", "fantasy", "wireframe", "black", "luxury", "dracula",
	"cmyk", "autumn", "business", "acid", "lemonade", "night", "coffee",
	"winter", "dim", "nord", "sunset", "caramellatte", "abyss", "silk",
}

func isValidDaisyUITheme(theme string) bool {
	for _, validTheme := range validDaisyUIThemes {
		if theme == validTheme {
			return true
		}
	}
	return false
}

func TestDarkTheme(t *testing.T) {
	if DarkTheme == "" {
		t.Error("DarkTheme should not be empty")
	}

	if !isValidDaisyUITheme(DarkTheme) {
		t.Errorf("DarkTheme %q is not a valid DaisyUI theme", DarkTheme)
	}
}

func TestLightTheme(t *testing.T) {
	if LightTheme == "" {
		t.Error("LightTheme should not be empty")
	}

	if !isValidDaisyUITheme(LightTheme) {
		t.Errorf("LightTheme %q is not a valid DaisyUI theme", LightTheme)
	}
}

func TestThemesAreDifferent(t *testing.T) {
	if DarkTheme == LightTheme {
		t.Errorf("DarkTheme and LightTheme should be different, both are %q", DarkTheme)
	}
}

func TestValidDaisyUIThemes(t *testing.T) {
	// Test that our helper function works correctly
	if !isValidDaisyUITheme("dark") {
		t.Error("Expected 'dark' to be a valid DaisyUI theme")
	}

	if !isValidDaisyUITheme("winter") {
		t.Error("Expected 'winter' to be a valid DaisyUI theme")
	}

	if isValidDaisyUITheme("invalid-theme") {
		t.Error("Expected 'invalid-theme' to be invalid")
	}

	if isValidDaisyUITheme("") {
		t.Error("Expected empty string to be invalid")
	}
}
