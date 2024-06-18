package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"features/features"

	"github.com/PuerkitoBio/goquery"
)

type FeatureResponse struct {
	Features []int `json:"features"`
}

// Extracts features from HTML content
func extractFeatures(html string) ([]int, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	// Call feature extraction functions from your custom package
	return []int{
		features.HasTitle(doc),
		features.HasInput(doc),
		features.HasButton(doc),
		features.HasImage(doc),
		features.HasSubmit(doc),
		features.HasLink(doc),
		features.HasPassword(doc),
		features.HasEmailInput(doc),
		features.HasHiddenElement(doc),
		features.HasAudio(doc),
		features.HasVideo(doc),
		features.NumberOfInputs(doc),
		features.NumberOfButtons(doc),
		features.NumberOfImages(doc),
		features.NumberOfOption(doc),
		features.NumberOfList(doc),
		features.NumberOfTH(doc),
		features.NumberOfTR(doc),
		features.NumberOfHref(doc),
		features.NumberOfParagraph(doc),
		features.NumberOfScript(doc),
		features.LengthOfTitle(doc),
		features.HasH1(doc),
		features.HasH2(doc),
		features.HasH3(doc),
		features.LengthOfText(doc),
		features.NumberOfClickableButton(doc),
		features.NumberOfA(doc),
		features.NumberOfImg(doc),
		features.NumberOfDiv(doc),
		features.NumberOfFigure(doc),
		features.HasFooter(doc),
		features.HasForm(doc),
		features.HasTextArea(doc),
		features.HasIFrame(doc),
		features.HasTextInput(doc),
		features.NumberOfMeta(doc),
		features.HasNav(doc),
		features.HasObject(doc),
		features.HasPicture(doc),
		features.NumberOfSources(doc),
		features.NumberOfSpan(doc),
		features.NumberOfTable(doc),
	}, nil
}

func featureExtractionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	htmlContent := r.FormValue("html")
	if htmlContent == "" {
		http.Error(w, "HTML content is required", http.StatusBadRequest)
		return
	}

	features, err := extractFeatures(htmlContent)
	if err != nil {
		http.Error(w, "Error extracting features", http.StatusInternalServerError)
		return
	}

	response := FeatureResponse{
		Features: features,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/extract_features", featureExtractionHandler)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
