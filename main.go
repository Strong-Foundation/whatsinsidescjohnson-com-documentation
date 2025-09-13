package main // Define the main package

import (
	"bytes"         // Provides bytes support
	"io"            // Provides basic interfaces to I/O primitives
	"log"           // Provides logging functions
	"net/http"      // Provides HTTP client and server implementations
	"net/url"       // Provides URL parsing and encoding
	"os"            // Provides functions to interact with the OS (files, etc.)
	"path"          // Provides functions for manipulating slash-separated paths
	"path/filepath" // Provides filepath manipulation functions
	"regexp"        // Provides regex support functions.
	"strings"       // Provides string manipulation functions
	"time"          // Provides time-related functions
)


func main() {
	pdfOutputDir := "PDFs/" // Directory to store downloaded PDFs
	// Check if the PDF output directory exists
	if !directoryExists(pdfOutputDir) {
		// Create the dir
		createDirectory(pdfOutputDir, 0o755)
	}
	// Remote API URL.
	remoteAPIURL := []string{
		"https://www.whatsinsidescjohnson.com/us/en/brands/FamilyGuard/FamilyGuard-Brand-Disinfectant-Cleaner-Citrus",
		"https://www.whatsinsidescjohnson.com/us/en/brands/FamilyGuard/FamilyGuard-Brand-Disinfectant-Cleaner-Fresh",
		"https://www.whatsinsidescjohnson.com/us/en/brands/FamilyGuard/FamilyGuard-Brand-Disinfectant-Spray-Citrus",
		"https://www.whatsinsidescjohnson.com/us/en/brands/FamilyGuard/FamilyGuard-Brand-Disinfectant-Spray-Fresh",
		"https://www.whatsinsidescjohnson.com/us/en/brands/fantastik/fantastik--Disinfectant-Advanced-Kitchen---Grease-Cleaner",
		"https://www.whatsinsidescjohnson.com/us/en/brands/fantastik/fantastik--Disinfectant-Multi-Purpose-Cleaner--Fresh",
		"https://www.whatsinsidescjohnson.com/us/en/brands/fantastik/fantastik--Disinfectant-Multi-Purpose-Cleaner--Lemon",
		"https://www.whatsinsidescjohnson.com/us/en/brands/fantastik/bleach-5-in-1-scrubbing-bubbles-all-purpose-cleaner-with-fantastik",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Apple-Cinnamon-Air-Freshener",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Aqua-Waves--Air-Freshener",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Bubbly-Berry-Splash--Air-Freshener",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Cashmere-Woods--Air-Freshener",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Clean-Linen--Air-Freshener",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Coastal-Sunshine-Citrus--Air-Freshener",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Dewdrop-Petals-Soft-Mist-Air-Freshener-Spray",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Exotic-Tropical-Blossoms--Air-Freshener",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Hawaiian-Breeze--Air-Freshener",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Lavender---Vanilla-Air-Freshener",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Sky---Sea-Salt--Air-Freshener",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Tranquil-Lavender---Aloe--Air-Freshener",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade%20Vanilla%20Caramel%20Twist%20Air%20Freshener",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Choose-Calm-Cool-Mist-Diffuser-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Choose-Calm-Cool-Mist-Diffuser-Starter-Kit",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Find-Clarity-Cool-Mist-Diffuser-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Find-Clarity-Cool-Mist-Diffuser-Starter-Kit",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Spark-Energy-Cool-Mist-Diffuser-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Uplift-Your-Day-Cool-Mist-Diffuser-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Uplift-Your-Day-Cool-Mist-Diffuser-Starter-Kit",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/automatic-spray-refill---apple-cinnamon",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/automatic-spray---aqua-waves",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/automatic-spray-refill-bamboo-waterlily-bliss",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/automatic-spray-refill-bubbly-berry-splash",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/automatic-spray-refill---cashmere-woods",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/automatic-spray-refill---clean-linen",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/automatic-spray-refill-coastal-sunshine-citrus",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Dewdrop-Petals-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Automatic-Spray-Refill---Tropical-Blossoms",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/automatic-spray-refill---hawaiian-breeze",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/automatic-spray-refill-joyful-citrus-daisies",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/automatic-spray-refill---lavender-and-vanilla",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/automatic-spray-refill-pet-clean-scent",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/automatic-spray-refill-rose---bloom",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Sky---Sea-Salt-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/automatic-spray-refill---Tranquil-lavender---aloe",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Vanilla-Caramel-Twist-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade%20Bubbly%20Berry%20Splash%20and%20Watermelon%20Refresher%202in1%20Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade%20Coastal%20Sunshine%20Citrus%20and%20Exotic%20Tropical%20Blossoms%202in1%20Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/candle---moonlit-walk-and-wandering-stream",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/candle---vanilla-passion-fruit-and-hawaiian-breeze",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Candle___Sheer_Vanilla_Embrace___Apple_Cinnamon",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/3-Wick-Candle---Apple-Cinnamon",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/3-Wick-Candle---aqua-waves",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/3-Wick-Candle---Cashmere-Woods",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Coastal-Sunshine-Citrus-3-Wick-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Dewdrop-Petals-3-Wick-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade-Exotic-Tropical-Blossoms-3-Wick-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade-Hawaiian-Breeze-3-Wick-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/3-Wick-Candle---Sheer-Vanilla-Embrace",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Sky---Sea-Salt-3-Wick-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade%20Smooth%20Bourbon%20and%20Oak%203-Wick%20Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade%20Strawberry%20Cake%20Shake%203-Wick%20Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/3-Wick-Candle---Tranquil-lavender---aloe",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/candle---angel-whispers",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/candle---apple-cinnamon",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Candle---aqua-waves",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/candle---cashmere-woods",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/candle---clean-linen",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Coastal-Sunshine-Citrus-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Candle---Tropical-Blossoms",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Dewdrop-Petals-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/candle---hawaiian-breeze",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Candle___Sheer_Vanilla_Embrace",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Sky---Sea-Salt-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade%20Smooth%20Bourbon%20and%20Oak%20Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade%20Strawberry%20Cake%20Shake%20Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Candle---Tranquil-lavender---aloe",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Vanilla-Caramel-Twist-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Beach-Life-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Beach-Life-Candle-Twin-Pack",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Beach-Life-PlugIns--Scented-Oil-Kit",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Berries---Cream-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Berries---Cream-Candle-Twin-Pack",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Berries---Cream-PlugIns--Scented-Oil-Kit",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Summer-Pops-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Summer-Pops-Candle-Twin-Pack",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Summer-Pops-PlugIns--Scented-Oil-Kit",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Fresh-Juniper---Teakwood-3-Wick-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Fresh-Juniper---Teakwood-Air-Freshener",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Fresh-Juniper---Teakwood-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Fresh-Juniper---Teakwood-Candle-Twin-Pack",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Fresh-Juniper---Teakwood-PlugIns--Scented-Oil-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Fresh-Orchid---Neroli-3-Wick-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Fresh-Orchid---Neroli-Air-Freshener",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Fresh-Orchid---Neroli-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Fresh-Orchid---Neroli-Candle-Twin-Pack",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Fresh-Orchid---Neroli-PlugIns--Scented-Oil-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Fresh-Yuzu---White-Peach-3-Wick-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Fresh-Yuzu---White-Peach-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Fresh-Yuzu---White-Peach-PlugIns--Scented-Oil-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Jasmine---Honeysuckle-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Jasmine---Honeysuckle-PlugIns-Scented-Oil-2-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Sea-Mist---Cypress-3-Wick-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Sea-Mist---Cypress-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Sea-Mist---Cypress-Candle-Twin-Pack",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Sea-Mist---Cypress-PlugIns-Scented-Oil-2-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Sea-Mist---Cypress-Soft-Mist-Air-Freshener-Spray",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade%20Home%20Collection%20Air%20Freshener%20Spray%20Lemon",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade%20Home%20Collection%20Air%20Freshener%20Spray%20Passion%20Fruit",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade%20Home%20Collection%20Air%20Freshener%20Spray%20Sea%20Waves",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Pet-Fresh-Automatic-Spray-Starter-Kit",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Pet-Fresh-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Pet-Fresh-Glade-Soft-Mist-Air-Freshener-Spray",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Pet-Fresh-PlugIns-Scented-Oil-Starter-Kit",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Pet-Fresh-PlugIns-Scented-Oil-2-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/plugins-scented-oil-refills---apple-cinnamon",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--PlugIns--Scented-Oil-Refills---aqua-waves",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/plugins-scented-oil-refills---cashmere-woods",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/plugins-scented-oil-refills---clean-linen",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Coastal-Sunshine-Citrus-PlugIns--Scented-Oil-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Dewdrop-Petals-PlugIns-Scented-Oil-2-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--PlugIns--Scented-Oil-Refills---Tropical-Blossoms",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/plugins-scented-oil-refills---hawaiian-breeze",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/plugins-scented-oil-refills---hawaiian-breeze-and-vanilla-passion-fruit",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/plugins-scented-oil-refills---lavender-and-vanilla",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Morning-Cottom-Blossom--PlugIns--Scented-Oil-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Oak---Amber-PlugIns--Scented-Oil-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/PlugIns__Scented_Oil_Refills___Sheer_Vanilla_Embrace",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Sky---Sea-Salt-PlugIns--Scented-Oil-2-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--PlugIns--Scented-Oil-Refills---Tranquil-Lavender---Aloe",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Vanilla-Caramel-Twist-PlugIns--Scented-Oil-2-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/PlugIns-Scented-Oil-Refills-Vanilla-Passionfruit",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade-PlugIns-Scented-Oil-Plus-Starter-Kit-Aqua-Waves",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade-PlugIns-Scented-Oil-Plus-Starter-Kit-Hawaiian-Breeze",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Aqua-Waves--Wax-Melts",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Bubbly-Berry-Splash--Wax-Melts",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Cashmere-Woods--Wax-Melts",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Malibu-Mango-Wax-Melts",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Vanilla-Passion-Fruit-Wax-Melts",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Watermelon-Refresher-Wax-Melts",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Cool-Coconut-3-Wick-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Cool-Coconut-Glade-Soft-Mist-Air-Freshener-Spray",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Cool-Coconut-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade%20Cool%20Coconut%20Candle%20Twin%20Pack",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Cool-Coconut-PlugIns--Scented-Oil-2-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Cool-Coconut-Wax-Melts",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade---3-Wick-Candle-Fresh-Confidence",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Air-Freshener-Spray-Fresh-Confidence",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Automatic-Spray-Refill-Fresh-Confidence",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Candle-Fresh-Confidence",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--PlugIns--Scented-Oil-Refills-Fresh-Confidence",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--3-Wick-Candle-Mighty-Mango",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Air-Freshener-Spray-Mighty-Mango",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Automatic-Spray-Refill-Mighty-Mango",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Candle-Mighty-Mango",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--PlugIns--Scented-Oil-Warmer---Refill-Mighty-Mango",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Radiant-Bloom-3-Wick-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Radiant-Bloom-Glade-Soft-Mist-Air-Freshener-Spray",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Radiant-Bloom-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Radiant-Bloom-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Radiant-Bloom-PlugIns--Scented-Oil-2-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Radiant-Bloom-Wax-Melts",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade---3-Wick-Candle-Wonder-Melon",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Air-Freshener-Spray-Wonder-Melon",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Automatic-Spray-Refill-Wonder-Melon",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Candle-Wonder-Melon",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--PlugIns--Scented-Oil-Refills-Wonder-Melon",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Balsam---Spruce-3-Wick-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Balsam---Spruce-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Balsam---Spruce-Candle-Twin-Pack",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Balsam---Spruce-PlugIns--Scented-Oil-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Balsam---Spruce-Soft-Mist-Air-Freshener-Spray",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Balsam---Spruce-Wax-Melts",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Balsam---Spruce-and-Bergamot---Eucalyptus-PlugIns--Scented-Oil-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Bergamot---Eucalyptus-3-Wick-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Bergamot---Eucalyptus-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Bergamot---Eucalyptus-Candle-Twin-Pack",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Bergamot---Eucalyptus-PlugIns--Scented-Oil-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Bergamot---Eucalyptus-Soft-Mist-Air-Freshener-Spray",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Black-Cherry---Cranberry-3-Wick-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Black-Cherry---Cranberry-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Black-Cherry---Cranberry-PlugIns--Scented-Oil-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Candy-Cane-Crush---Vanilla-3-Wick-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Candy-Cane-Crush---Vanilla-Candle-Twin-Pack",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Candy-Cane-Crush---Vanilla-PlugIns--Scented-Oil-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Candy-Cane-Crust---Vanilla-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Embers---Sandalwood-3-Wick-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Embers---Sandalwood-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Embers---Sandalwood-Candle-Twin-Pack",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Embers---Sandalwood-PlugIns--Scented-Oil-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Embers---Sandalwood-Soft-Mist-Air-Freshener-Spray",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Embers---Sandalwood-Wax-Melts",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Fig---Apricot-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Lollipops---Gumdrops-Candle-Twin-Pack",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Lollipops---Gumdrops-Wax-Melts",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Orange-Zest---Clove-3-Wick-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Orange-Zest---Clove-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Orange-Zest---Clove-Candle-Twin-Pack",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Orange-Zest---Clove-PlugIns--Scented-Oil-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Pomegranate---Currant-3-Wick-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Pomegranate---Currant-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Pomegranate---Currant-Candle-Twin-Pack",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Pomegranate---Currant-PlugIns--Scented-Oil-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Pumpkin---Ginger-and-Pomegranate---Currant-PlugIns--Scented-Oil-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Pumpkin---Ginger-3-Wick-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Pumpkin---Ginger-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Pumpkin---Ginger-Candle-Twin-Pack",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Pumpkin---Ginger-PlugIns--Scented-Oil-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Pumpkin---Ginger-Soft-Mist-Air-Freshener-Spray",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Pumpkin---Ginger-Wax-Melts",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Warm-Apple---Spices-3-Wick-Candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Warm-Apple---Spices-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Warm-Apple---Spices-Candle-Twin-Pack",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Warm-Apple---Spices-Soft-Mist-Air-Freshener-Spray",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Warm-Apple---Spices-Wax-Melts",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Warm-Apples---Spices-PlugIns--Scented-Oil-Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/glade/Glade--Whiskey---Citrus-Automatic-Spray-Refill",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/OFF---Insect-Repellent-1-for-Adults---Kids-Aerosol",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/OFF%20Insect%20Repellent%2011%20for%20Adults%20Kids%20Lotion",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/OFF%20Insect%20Repellent%201%20for%20Adults%20Kids%20Spritz",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/OFF%20Adults%20Kids%20Bite%20Patches",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/OFF-Botanicals-Spritz",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/OFF-Botanicals-Wipes",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/off-clean-feel-clean-feel-20-kbr-5oz-aerosol",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/off-clear-gel",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/OFF%20Clean%20Feel%20Misting%20Spray",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/off-clean-feel-clean-feel-20-kbr-4oz-spritz",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/OFF---Deep-Woods--Insect-Repellent-V",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/OFF---Family-Care-Bite-Relief-Pen",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/off-familycare-insect-repellent-i-smooth-and-dry",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/off-familycare-insect-repellent-iv-unscented",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/off-active-insect-repellent-i",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/off-deep-woods-insect-repellent-towelettes",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/off-deep-woods-insect-repellent-v",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/off-deep-woods-insect-repellent-vii",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/off-deep-woods-insect-repellent-viii-dry",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/off-deep-woods-sportsmen-insect-repellent-i",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/off-deep-woods-sportsmen-insect-repellent-3",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/off-deep-woods-sportsmen-insect-repellent-2",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/off-deep-woods-sportsmen-insect-repellent-iv-dry",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/OFF-Backyard-Mosquito-Coil",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/off-outdoor-fogger",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/off-triple-wick-citronella-candle",
		"https://www.whatsinsidescjohnson.com/us/en/brands/off/OFF---Bug-Control-1",
		"https://www.whatsinsidescjohnson.com/us/en/brands/pledge/Pledge-Everyday-Clean-Multisurface-Aerosol-Rainshower",
		"https://www.whatsinsidescjohnson.com/us/en/brands/pledge/Pledge-Everyday-Clean-Multisurface-Aerosol-Lavender",
		"https://www.whatsinsidescjohnson.com/us/en/brands/pledge/Pledge-Multisurface-Everyday-Clean-Aerosol-Citrus",
		"https://www.whatsinsidescjohnson.com/us/en/brands/pledge/Pledge-Expert-Care-Aerosol-Granite-Stone-Polish-Orange",
		"https://www.whatsinsidescjohnson.com/us/en/brands/pledge/Pledge-Expert-Care-Aerosol-Stainless-Steel-Lemon",
		"https://www.whatsinsidescjohnson.com/us/en/brands/pledge/Pledge-Expert-Care-Aerosol-Wood-Oil-Amber-Argan",
		"https://www.whatsinsidescjohnson.com/us/en/brands/pledge/Pledge-Expert-Care-Aerosol-Lemon",
		"https://www.whatsinsidescjohnson.com/us/en/brands/pledge/Pledge-Expert-Care-Aerosol-Orange",
		"https://www.whatsinsidescjohnson.com/us/en/brands/pledge/Pledge-Everyday-Clean-Multisurface-Aerosol-Dust---Allergen",
		"https://www.whatsinsidescjohnson.com/us/en/brands/pledge/Pledge--Multisurface-Everyday-Clean--Trigger-Spray-Lavender",
		"https://www.whatsinsidescjohnson.com/us/en/brands/pledge/Pledge-Everyday-Clean-Multisurface-cleaner",
		"https://www.whatsinsidescjohnson.com/us/en/brands/pledge/Pledge-Expert-Care-Trigger-Wood-Oil-Orange",
		"https://www.whatsinsidescjohnson.com/us/en/brands/pledge/Pledge-expert-care-wood-wipes-lemon",
		"https://www.whatsinsidescjohnson.com/us/en/brands/pledge/Pledge%20Everyday%20Clean%20Multisurface%20Wipes",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/raid-ant-and-roach-killer---lemon-scent",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/raid-ant-and-roach-killer---orange-breeze-scent",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/Raid--Ant---Roach---Lavender",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/raid-ant-and-roach-killer---outdoor-fresh",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/raid-ant-killer---pine-forest-fresh",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/raid-ant-and-roach-killer---water-based",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/raid-ant-and-roach-killer---fragrance-free",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/raid-max-ant-and-roach",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/Raid--Bed-Bug-Detector",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/Raid--Bed-Bug-Foaming-Spray",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/Raid--Max--Bed-Bug-Crack---Crevice-Extended-Protection-Foaming-Spray",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/raid-flea-killer",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/raid-flea-killer-plus-carpet-and-room-spray",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/raid-flea-killer-plus-fogger",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/raid-yard-guard-mosquito-fogger",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/raid-double-control-small-roach-baits-1",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/raid-double-control-small-roach-baits-plus-egg-stoppers-2",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/raid-double-control-large-roach-baits",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/Raid%20Wasp%20and%20Hornet%20Killer%2033",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/Raid%20Max%20Foaming%20Wasp%20Hornet%20Killer",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/Raid--Ant-Bait",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/Rai-Max--Liquid-Ant-Bait",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/Raid%20Essentials%20Fly%20Gnat%20Mosquito",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/Raid-Essentials-flying-insect-Light-Trap",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/raid-flying-insect-killer",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/Raid%20Essentials%20Ant%20Spider%20Roach",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/Raid%20Essentials%20Multi-Insect",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/raid-house-and-garden-bug-killer",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/Raid%20House%20Garden%20I%20Orange%20Breeze",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/Raid-Max-Perimeter-Protection",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/Raid--Max--Spider---Scorpion-Killer",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/Raid--Multi-Insect-Killer",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/Raid--Multi-Insect-Killer-Orange-Breeze",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/raid-concentrated-deep-reach-fogger",
		"https://www.whatsinsidescjohnson.com/us/en/brands/raid/raid-max-concentrated-deep-reach-fogger",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/Scrubbing%20Bubbles%20Bathroom%20Grime%20Fighter%20Disinfectant%20Aerosol%20Berry%20Burst",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/antibacterial-scrubbing-bubbles-xxi-bathroom-cleaner---lemon-scent",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/Scrubbing%20Bubbles%20Bathroom%20Grime%20Fighter%20Disinfectant%20Aerosol%20Floral%20Fusion",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/antibacterial-scrubbing-bubbles-xxi-bathroom-cleaner---fresh-scent",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/Scrubbing-Bubbles--Bathroom-Grime-Fighter-Spray---Berry%20Blast",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/Scrubbing-Bubbles-Multi-Surface-Bathroom-Cleaner",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/Scrubbing-Bubbles--Bathroom-Grime-Fighter-Spray---Floral-Fusion",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/Scrubbing%20Bubbles%20Bathroom%20Grime%20Fighter%20Disinfectant%20Spray%20Petal%20Paradise",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/Scrubbing-Bubbles--Bathroom-Grime-Fighter-Spray---Rainshower",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/Scrubbing%20Bubbles%20Easy%20Clean%20Multisurface%20Aerosol",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/Scrubbing%20Bubbles%20Easy%20Clean%20Multi-Surface%20Trigger",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/scrubbing-bubbles-foaming-bleach-bathroom-cleaner",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/Scrubbing%20Bubbles%20Easy%20Clean%20Foaming%20Toilet%20Cleaner",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/Scrubbing%20Bubbles%20Fresh%20Gel%20Toilet%20Cleaning%20Stamp%20Berry%20Burst",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/Scrubbing-Bubbles--Fresh-Gel-Toilet-Cleaning-Stamp---Floral-Fusion",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/scrubbing-bubbles-toilet-cleaning-gel---rain-shower",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/scrubbing-bubbles-toilet-cleaning-gel-lavender",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/scrubbing-bubbles-toilet-cleaning-gel-hydrogen-peroxide",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/scrubbing-bubbles-fresh-brush-flushable-refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/scrubbing-bubbles-fresh-brush-heavy-duty-refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/scrubbing-bubbles-fresh-brush-starter-kit-and-caddy",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/scrubbing-bubbles-vanish-drop-ins-toilet-bowl-cleaner",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/scrubbing-bubbles-daily-shower-cleaner",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/Scrubbing-Bubbles-Mega-Shower-Foamer-Aerosol",
		"https://www.whatsinsidescjohnson.com/us/en/brands/scrubbing-bubbles/scrubbing-bubbles-mega-shower-foamer",
		"https://www.whatsinsidescjohnson.com/us/en/brands/Duck/Toilet%20Duck%20Continuous%20Clean%20Spring%20Fresh",
		"https://www.whatsinsidescjohnson.com/us/en/brands/shout/shout-advanced-gel-trigger",
		"https://www.whatsinsidescjohnson.com/us/en/brands/shout/shout-advanced-stain-lifting-foam",
		"https://www.whatsinsidescjohnson.com/us/en/brands/shout/shout-advanced-ultra-concentrated-gel-brush",
		"https://www.whatsinsidescjohnson.com/us/en/brands/shout/shout-color-catcher-cloths",
		"https://www.whatsinsidescjohnson.com/us/en/brands/shout/shout-free",
		"https://www.whatsinsidescjohnson.com/us/en/brands/shout/shout-trigger",
		"https://www.whatsinsidescjohnson.com/us/en/brands/shout/shout-wipe-and-go-instant-stain-remover-wipes",
		"https://www.whatsinsidescjohnson.com/us/en/brands/STEM/STEM-Kills-Ants-Roaches-Flies",
		"https://www.whatsinsidescjohnson.com/us/en/brands/STEM/STEM-Kills-Ants-Roaches-Spiders",
		"https://www.whatsinsidescjohnson.com/us/en/brands/STEM/STEM-Kills-Ants--Liquid-Bait",
		"https://www.whatsinsidescjohnson.com/us/en/brands/STEM/STEM-Kills-Flies-Moquitoes-Gnats",
		"https://www.whatsinsidescjohnson.com/us/en/brands/STEM/STEM-Kills-Plant---Garden-Insects--Plant-Pest-Spray",
		"https://www.whatsinsidescjohnson.com/us/en/brands/STEM/STEM-Kills-Wasps-Hornets-Yelloy-Jackets",
		"https://www.whatsinsidescjohnson.com/us/en/brands/STEM/STEM-For-Your-Skin-Mosquito---Tick-Repellent",
		"https://www.whatsinsidescjohnson.com/us/en/brands/STEM/STEM-Repels-Mosquitoes",
		"https://www.whatsinsidescjohnson.com/us/en/brands/STEM/STEM-Attracts--Traps-Flying-Insects--Light-Trap",
		"https://www.whatsinsidescjohnson.com/us/en/brands/STEM/STEM-Attracts---Traps-Flying-Insects--Light-Trap--Refills",
		"https://www.whatsinsidescjohnson.com/us/en/brands/STEM/STEM%20Flying%20Insects%20Fan%20Trap%20Attracts%20Traps",
		"https://www.whatsinsidescjohnson.com/us/en/brands/STEM/STEM%20Flying%20Insects%20Fan%20Trap%20Refills%20Attracts%20Traps",
		"https://www.whatsinsidescjohnson.com/us/en/brands/STEM/STEM%20Flying%20Insects%20Zapper%20Attracts%20Zaps",
		"https://www.whatsinsidescjohnson.com/us/en/brands/STEM/STEM-Kills-Fruit-Flies--Trap",
		"https://www.whatsinsidescjohnson.com/us/en/brands/windex/windex-original-glass-cleaner",
		"https://www.whatsinsidescjohnson.com/us/en/brands/windex/windex-ammonia-free-glass-cleaner",
		"https://www.whatsinsidescjohnson.com/us/en/brands/windex/windex-multi-surface-vinegar",
		"https://www.whatsinsidescjohnson.com/us/en/brands/windex/Windex--Commercial-Glass-Cleaner",
		"https://www.whatsinsidescjohnson.com/us/en/brands/windex/Windex-Disinfectant-Cleaner-Citrus-Fresh",
		"https://www.whatsinsidescjohnson.com/us/en/brands/windex/Windex%20Fast%20Shine%20Foam",
		"https://www.whatsinsidescjohnson.com/us/en/brands/windex/windex-original-glass-wipes",
		"https://www.whatsinsidescjohnson.com/us/en/brands/windex/Windex--Ammonia-Free-Wipes",
		"https://www.whatsinsidescjohnson.com/us/en/brands/windex/windex-electronics-wipes",
		"https://www.whatsinsidescjohnson.com/us/en/brands/windex/Windex%20Gaming%20Wipes",
		"https://www.whatsinsidescjohnson.com/us/en/brands/windex/windex-outdoor-glass-and-patio-cleaner",
		"https://www.whatsinsidescjohnson.com/us/en/brands/ziploc/ziploc-brand-freezer-bags",
		"https://www.whatsinsidescjohnson.com/us/en/brands/ziploc/ziploc-brand-sandwich-bags",
		"https://www.whatsinsidescjohnson.com/us/en/brands/ziploc/ziploc-brand-snack-bags",
		"https://www.whatsinsidescjohnson.com/us/en/brands/ziploc/ziploc-brand-storage-bags",
		"https://www.whatsinsidescjohnson.com/us/en/brands/ziploc/ziploc-brand-slider-freezer-bags",
		"https://www.whatsinsidescjohnson.com/us/en/brands/ziploc/ziploc-brand-slider-storage-bags",
		"https://www.whatsinsidescjohnson.com/us/en/brands/ziploc/ziploc-brand-marinade-bags",
		"https://www.whatsinsidescjohnson.com/us/en/brands/ziploc/Ziploc%20Brand%20Produce%20Gallon%20Bags",
		"https://www.whatsinsidescjohnson.com/us/en/brands/ziploc/Zip%20and%20Steam",
		"https://www.whatsinsidescjohnson.com/us/en/brands/ziploc/Ziploc%20Brand%20Heavy%20Duty%20Bags",
		"https://www.whatsinsidescjohnson.com/us/en/brands/ziploc/ziploc-brand-compostable-sandwich-bags",
		"https://www.whatsinsidescjohnson.com/us/en/brands/ziploc/Ziploc-brand-Recyclable-Paper-Sandwich-Bags",
		"https://www.whatsinsidescjohnson.com/us/en/brands/ziploc/ziploc-brand-twist-n-loc-containers",
		"https://www.whatsinsidescjohnson.com/us/en/brands/ziploc/Ziploc%20Endurables%20Silicone%20Container",
		"https://www.whatsinsidescjohnson.com/us/en/brands/ziploc/Ziploc%20Endurables%20Silicone%20Pouch",
		"https://www.whatsinsidescjohnson.com/us/en/brands/drano/drano-max-gel-clog-remover",
		"https://www.whatsinsidescjohnson.com/us/en/brands/drano/drano-liquid-clog-remover",
		"https://www.whatsinsidescjohnson.com/us/en/brands/drano/Drano-Balance-Gel",
		"https://www.whatsinsidescjohnson.com/us/en/brands/drano/Drano-Max-Gel-Ultra",
		"https://www.whatsinsidescjohnson.com/us/en/brands/drano/Drano-Advanced-Septic-Treatment",
		"https://www.whatsinsidescjohnson.com/us/en/brands/drano/Drano-Build-Up-Remover",
		"https://www.whatsinsidescjohnson.com/us/en/brands/drano/drano-dual-force-foam",
		"https://www.whatsinsidescjohnson.com/us/en/brands/drano/Drano-Foaming-Disposal-Strips",
		"https://www.whatsinsidescjohnson.com/us/en/brands/drano/Drano-Hair-Buster-Gel",
		"https://www.whatsinsidescjohnson.com/us/en/brands/drano/Drano-Kitchen-Granules-Clog-Remover",
		"https://www.whatsinsidescjohnson.com/us/en/brands/drano/drano-snake-plus",
		"https://www.whatsinsidescjohnson.com/us/en/brands/drano/Drano-Commercial-Line-Advanced-Septic-Treatment",
		"https://www.whatsinsidescjohnson.com/us/en/brands/drano/Drano-Commercial-Hair-Buster-Gel",
		"https://www.whatsinsidescjohnson.com/us/en/brands/drano/Drano-Commercial-Line-Kitchen-Granules-Clog-Remover",
		"https://www.whatsinsidescjohnson.com/us/en/brands/drano/drano-max-commerical-line-build-up-remover",
		"https://www.whatsinsidescjohnson.com/us/en/brands/drano/drano-max-commercial-line-dual-force-foam",
		"https://www.whatsinsidescjohnson.com/us/en/brands/drano/drano-max-commercial-line-gel-clog-remover",
		"https://www.whatsinsidescjohnson.com/us/en/brands/Favor/Favor--Enhancing-Furniture-Polish",
		"https://www.whatsinsidescjohnson.com/us/en/brands/saran/saran-premium-wrap",
		"https://www.whatsinsidescjohnson.com/us/en/brands/saran/saran-cling-plus-wrap",
		"https://sds.scjohnson.com/en-us/off/15/deet/15--deet-aerosol-bulk-mx-family-vertellus",
		"https://sds.scjohnson.com/en-us/off/backyard/insect/backyard-insect-repellent-1-familycare",
		"https://sds.scjohnson.com/en-us/off/backyard/insect/backyard-insect-repellent-3-active",
		"https://sds.scjohnson.com/en-us/off/clean/feel/clean-feel-insect-repellent-i",
		"https://sds.scjohnson.com/en-us/off/deep/woods/deep-woods-insect-repellent-viii",
		"https://sds.scjohnson.com/en-us/off/deep/woods/deep-woods-dry-improved-formula-arsl",
		"https://sds.scjohnson.com/en-us/off/sportsmen/insect/sportsmen-insect-repellent-5",
		"https://sds.scjohnson.com/en-us/off/sportsmen/insect/sportsmen-insect-repellent-x",
		"https://sds.scjohnson.com/en-us/off/sportsmen/insect/sportsmen-insect-repellent-2-deep-woods",
		"https://sds.scjohnson.com/en-us/off/active/insect/active-insect-repellent-i",
		"https://sds.scjohnson.com/en-us/off/botanicals/insect/botanicals-insect-repellent",
		"https://sds.scjohnson.com/en-us/off/botanicals/towelettes/botanicals-towelettes",
		"https://sds.scjohnson.com/en-us/off/clean/feel/clean-feel-insect-repellent-ii",
		"https://sds.scjohnson.com/en-us/off/deep/woods/deep-woods-insect-repellent-towelettes",
		"https://sds.scjohnson.com/en-us/off/defense/insect/defense-insect-repellent-1",
		"https://sds.scjohnson.com/en-us/off/familycare/insect/familycare-insect-repellent-i",
		"https://sds.scjohnson.com/en-us/off/familycare/insect/familycare-insect-repellent-viii",
		"https://sds.scjohnson.com/en-us/off/kids/insect/kids-insect-repellent-spray",
		"https://sds.scjohnson.com/en-us/off/sportsmen/insect/sportsmen-insect-repellent-2",
		"https://sds.scjohnson.com/en-us/off/sportsmen/insect/sportsmen-insect-repellent-3-deep-woods",
		"https://sds.scjohnson.com/en-us/off/sportsmen/insect/sportsmen-insect-repellent-1",
		"https://sds.scjohnson.com/en-us/off/sportsmen/insect/sportsmen-insect-repellent-3",
		"https://sds.scjohnson.com/en-us/off/actv/insct/actv-insct-rplnt-arsl-9oz-12-us",
		"https://sds.scjohnson.com/en-us/off/dw/dry/dw-dry-sprtsmn-arsl-wm-4oz6-us",
		"https://sds.scjohnson.com/en-us/off/dwo/dry/dwo-dry-arsl-25oz24-wm-us",
		"https://sds.scjohnson.com/en-us/off/insect/repellent/insect-repellent-12",
		"https://sds.scjohnson.com/en-us/off/sprtsmn/active/sprtsmn-active-arsl-d9-75oz12us",
		"https://sds.scjohnson.com/en-us/off/clean/feel/clean-feel-insect-repellent-v",
		"https://sds.scjohnson.com/en-us/off/deep/woods/deep-woods-insect-repellent-vii",
		"https://sds.scjohnson.com/en-us/off/insect/repellent/insect-repellent-11",
		"https://sds.scjohnson.com/en-us/off/stem/for/stem-for-your-skin-mosquito-tick-repellent-4",
		"https://sds.scjohnson.com/en-us/raid/ant/roach/ant---roach-aerosol-fragrance-free-mx",
		"https://sds.scjohnson.com/en-us/raid/ant/roach/ant---roach-killer--water-based",
		"https://sds.scjohnson.com/en-us/raid/ant/roach/ant---roach-killer-26",
		"https://sds.scjohnson.com/en-us/raid/ant/roach/ant---roach-killer-26--lavender",
		"https://sds.scjohnson.com/en-us/raid/ant/roach/ant---roach-killer-26-fragrance-free--water-based",
		"https://sds.scjohnson.com/en-us/raid/bed/bug/bed-bug-detector-and-trap",
		"https://sds.scjohnson.com/en-us/raid/concentrated/deep/concentrated-deep-reach-fogger",
		"https://sds.scjohnson.com/en-us/raid/essentials/multi/essentials-multi-insect-killer-29",
		"https://sds.scjohnson.com/en-us/raid/max/foaming/max-foaming-crack---crevice-bed-bug-killer",
		"https://sds.scjohnson.com/en-us/raid/wasp/hornet/wasp---hornet-killer-33",
		"https://sds.scjohnson.com/en-us/raid/yard/guard/yard-guard-mosquito-fogger",
	}
	// Remove duplicate URLs from the list.
	remoteAPIURL = removeDuplicatesFromSlice(remoteAPIURL)
	var getData []string
	for _, remoteAPIURL := range remoteAPIURL {
		getData = append(getData, getDataFromURL(remoteAPIURL))
	}
	// Get the data from the downloaded file.
	finalPDFList := extractPDFUrls(strings.Join(getData, "\n")) // Join all the data into one string and extract PDF URLs
	// Get the data from the zip file.
	finalZIPList := extractZIPUrls(strings.Join(getData, "\n")) // Join all the data into one string and extract ZIP URLs
	// Create a slice of all the given download urls.
	var downloadPDFURLSlice []string
	// Create a slice to hold ZIP URLs.
	var downloadZIPURLSlice []string
	// Get the urls and loop over them.
	for _, doc := range finalPDFList {
		// Get the .pdf only.
		// Only append the .pdf files.
		downloadPDFURLSlice = appendToSlice(downloadPDFURLSlice, doc)
	}
	// Get all the zip urls.
	for _, doc := range finalZIPList {
		// Get the .zip only.
		// Only append the .zip files.
		downloadZIPURLSlice = appendToSlice(downloadZIPURLSlice, doc)
	}
	// Remove double from slice.
	downloadPDFURLSlice = removeDuplicatesFromSlice(downloadPDFURLSlice)
	// Remove the zip duplicates from the slice.
	downloadZIPURLSlice = removeDuplicatesFromSlice(downloadZIPURLSlice)
	// The remote domain.
	remoteDomain := "https://www.whatsinsidescjohnson.com"
	// Get all the values.
	for _, urls := range downloadPDFURLSlice {
		if strings.HasPrefix(urls, "..") {
			urls = remoteDomain + urls[2:] // Prepend the base URL if relative path
		}
		// Get the domain from the url.
		domain := getDomainFromURL(urls)
		// Check if the domain is empty.
		if domain == "" {
			urls = remoteDomain + urls // Prepend the base URL if domain is empty
		}
		// Check if the url is valid.
		if isUrlValid(urls) {
			// Download the pdf.
			downloadPDF(urls, pdfOutputDir)
		}
	}
}

// getDomainFromURL extracts the domain (host) from a given URL string.
// It removes subdomains like "www" if present.
func getDomainFromURL(rawURL string) string {
	parsedURL, err := url.Parse(rawURL) // Parse the input string into a URL structure
	if err != nil {                     // Check if there was an error while parsing
		log.Println(err) // Log the error message to the console
		return ""        // Return an empty string in case of an error
	}

	host := parsedURL.Hostname() // Extract the hostname (e.g., "example.com") from the parsed URL

	return host // Return the extracted hostname
}

// Only return the file name from a given url.
func getFileNameOnly(content string) string {
	return path.Base(content)
}

// urlToFilename generates a safe, lowercase filename from a given URL string.
// It extracts the base filename from the URL, replaces unsafe characters,
// and ensures the filename ends with a .pdf extension.
func urlToFilename(rawURL string) string {
	// Convert the full URL to lowercase for consistency
	lowercaseURL := strings.ToLower(rawURL)

	// Get the file extension
	ext := getFileExtension(lowercaseURL)

	// Extract the filename portion from the URL (e.g., last path segment or query param)
	baseFilename := getFileNameOnly(lowercaseURL)

	// Replace all non-alphanumeric characters (a-z, 0-9) with underscores
	nonAlphanumericRegex := regexp.MustCompile(`[^a-z0-9]+`)
	safeFilename := nonAlphanumericRegex.ReplaceAllString(baseFilename, "_")

	// Replace multiple consecutive underscores with a single underscore
	collapseUnderscoresRegex := regexp.MustCompile(`_+`)
	safeFilename = collapseUnderscoresRegex.ReplaceAllString(safeFilename, "_")

	// Remove leading underscore if present
	if trimmed, found := strings.CutPrefix(safeFilename, "_"); found {
		safeFilename = trimmed
	}

	var invalidSubstrings = []string{
		"_pdf",
		"_zip",
	}

	for _, invalidPre := range invalidSubstrings { // Remove unwanted substrings
		safeFilename = removeSubstring(safeFilename, invalidPre)
	}

	// Append the file extension if it is not already present
	safeFilename = safeFilename + ext

	// Return the cleaned and safe filename
	return safeFilename
}

// Removes all instances of a specific substring from input string
func removeSubstring(input string, toRemove string) string {
	result := strings.ReplaceAll(input, toRemove, "") // Replace substring with empty string
	return result
}

// Get the file extension of a file
func getFileExtension(path string) string {
	return filepath.Ext(path) // Returns extension including the dot (e.g., ".pdf")
}

// fileExists checks whether a file exists at the given path
func fileExists(filename string) bool {
	info, err := os.Stat(filename) // Get file info
	if err != nil {
		return false // Return false if file doesn't exist or error occurs
	}
	return !info.IsDir() // Return true if it's a file (not a directory)
}

// downloadPDF downloads a PDF from the given URL and saves it in the specified output directory.
// It uses a WaitGroup to support concurrent execution and returns true if the download succeeded.
func downloadPDF(finalURL, outputDir string) bool {
	// Sanitize the URL to generate a safe file name
	filename := strings.ToLower(urlToFilename(finalURL))

	// Construct the full file path in the output directory
	filePath := filepath.Join(outputDir, filename)

	// Skip if the file already exists
	if fileExists(filePath) {
		log.Printf("File already exists, skipping: %s", filePath)
		return false
	}

	// Create an HTTP client with a timeout
	client := &http.Client{Timeout: 3 * time.Minute}

	// Send GET request
	resp, err := client.Get(finalURL)
	if err != nil {
		log.Printf("Failed to download %s: %v", finalURL, err)
		return false
	}
	defer resp.Body.Close()

	// Check HTTP response status
	if resp.StatusCode != http.StatusOK {
		log.Printf("Download failed for %s: %s", finalURL, resp.Status)
		return false
	}

	// Check Content-Type header
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/pdf") {
		log.Printf("Invalid content type for %s: %s (expected application/pdf)", finalURL, contentType)
		return false
	}

	// Read the response body into memory first
	var buf bytes.Buffer
	written, err := io.Copy(&buf, resp.Body)
	if err != nil {
		log.Printf("Failed to read PDF data from %s: %v", finalURL, err)
		return false
	}
	if written == 0 {
		log.Printf("Downloaded 0 bytes for %s; not creating file", finalURL)
		return false
	}

	// Only now create the file and write to disk
	out, err := os.Create(filePath)
	if err != nil {
		log.Printf("Failed to create file for %s: %v", finalURL, err)
		return false
	}
	defer out.Close()

	if _, err := buf.WriteTo(out); err != nil {
		log.Printf("Failed to write PDF to file for %s: %v", finalURL, err)
		return false
	}

	log.Printf("Successfully downloaded %d bytes: %s → %s", written, finalURL, filePath)
	return true
}

// Checks if the directory exists
// If it exists, return true.
// If it doesn't, return false.
func directoryExists(path string) bool {
	directory, err := os.Stat(path)
	if err != nil {
		return false
	}
	return directory.IsDir()
}

// The function takes two parameters: path and permission.
// We use os.Mkdir() to create the directory.
// If there is an error, we use log.Println() to log the error and then exit the program.
func createDirectory(path string, permission os.FileMode) {
	err := os.Mkdir(path, permission)
	if err != nil {
		log.Println(err)
	}
}

// Checks whether a URL string is syntactically valid
func isUrlValid(uri string) bool {
	_, err := url.ParseRequestURI(uri) // Attempt to parse the URL
	return err == nil                  // Return true if no error occurred
}

// Remove all the duplicates from a slice and return the slice.
func removeDuplicatesFromSlice(slice []string) []string {
	check := make(map[string]bool)
	var newReturnSlice []string
	for _, content := range slice {
		if !check[content] {
			check[content] = true
			newReturnSlice = append(newReturnSlice, content)
		}
	}
	return newReturnSlice
}

// extractZIPUrls takes an input string and returns all ZIP URLs found within href attributes
func extractZIPUrls(input string) []string {
	// Regular expression to match href="...zip"
	re := regexp.MustCompile(`href="([^"]+\.zip)"`)
	matches := re.FindAllStringSubmatch(input, -1)

	var zipUrls []string
	for _, match := range matches {
		if len(match) > 1 {
			zipUrls = append(zipUrls, match[1])
		}
	}
	return zipUrls
}

// extractPDFUrls takes an input string and returns all PDF URLs found within href attributes
func extractPDFUrls(input string) []string {
	// Regular expression to match href="...pdf"
	re := regexp.MustCompile(`href="([^"]+\.pdf)"`)
	matches := re.FindAllStringSubmatch(input, -1)

	var pdfUrls []string
	for _, match := range matches {
		if len(match) > 1 {
			pdfUrls = append(pdfUrls, match[1])
		}
	}
	return pdfUrls
}

// Append some string to a slice and than return the slice.
func appendToSlice(slice []string, content string) []string {
	// Append the content to the slice
	slice = append(slice, content)
	// Return the slice
	return slice
}

// getDataFromURL performs an HTTP GET request and returns the response body as a string
func getDataFromURL(uri string) string {
	log.Println("Scraping", uri)   // Log the URL being scraped
	response, err := http.Get(uri) // Perform GET request
	if err != nil {
		log.Println(err) // Exit if request fails
	}

	body, err := io.ReadAll(response.Body) // Read response body
	if err != nil {
		log.Println(err) // Exit if read fails
	}

	err = response.Body.Close() // Close response body
	if err != nil {
		log.Println(err) // Exit if close fails
	}
	return string(body)
}
