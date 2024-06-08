package features

import (
	"github.com/PuerkitoBio/goquery"
)

// HasTitle checks if the doc has a title element
func HasTitle(doc *goquery.Document) int {
	if doc.Find("title").Length() == 0 {
		return 0
	}
	return 1
}

// HasInput checks if the doc has any input elements
func HasInput(doc *goquery.Document) int {
	if doc.Find("input").Length() > 0 {
		return 1
	}
	return 0
}

// HasButton checks if the doc has any button elements
func HasButton(doc *goquery.Document) int {
	if doc.Find("button").Length() > 0 {
		return 1
	}
	return 0
}

// HasImage checks if the doc has any image elements
func HasImage(doc *goquery.Document) int {
	if doc.Find("img").Length() > 0 {
		return 1
	}
	return 0
}

// HasSubmit checks if the doc has any submit input elements
func HasSubmit(doc *goquery.Document) int {
	if doc.Find("input[type=submit]").Length() > 0 {
		return 1
	}
	return 0
}

// HasLink checks if the doc has any link elements
func HasLink(doc *goquery.Document) int {
	if doc.Find("link").Length() > 0 {
		return 1
	}
	return 0
}

// HasPassword checks if the doc has any password input elements
func HasPassword(doc *goquery.Document) int {
	if doc.Find("input[type=password]").Length() > 0 {
		return 1
	}
	return 0
}

// HasEmailInput checks if the doc has any email input elements
func HasEmailInput(doc *goquery.Document) int {
	if doc.Find("input[type=email]").Length() > 0 {
		return 1
	}
	return 0
}

// HasHiddenElement checks if the doc has any hidden input elements
func HasHiddenElement(doc *goquery.Document) int {
	if doc.Find("input[type=hidden]").Length() > 0 {
		return 1
	}
	return 0
}

// HasAudio checks if the doc has any audio elements
func HasAudio(doc *goquery.Document) int {
	if doc.Find("audio").Length() > 0 {
		return 1
	}
	return 0
}

// HasVideo checks if the doc has any video elements
func HasVideo(doc *goquery.Document) int {
	if doc.Find("video").Length() > 0 {
		return 1
	}
	return 0
}

// NumberOfInputs returns the number of input elements in the doc
func NumberOfInputs(doc *goquery.Document) int {
	return doc.Find("input").Length()
}

// NumberOfButtons returns the number of button elements in the doc
func NumberOfButtons(doc *goquery.Document) int {
	return doc.Find("button").Length()
}

// NumberOfImages returns the number of image elements in the doc
func NumberOfImages(doc *goquery.Document) int {
	return doc.Find("img").Length()
}

// NumberOfOptions returns the number of option elements in the doc
func NumberOfOption(doc *goquery.Document) int {
	return doc.Find("option").Length()
}

// NumberOfLists returns the number of list elements in the doc
func NumberOfList(doc *goquery.Document) int {
	return doc.Find("li").Length()
}

// NumberOfTHs returns the number of table header elements (TH) in the doc
func NumberOfTH(doc *goquery.Document) int {
	return doc.Find("th").Length()
}

// NumberOfTRs returns the number of table row elements (TR) in the doc
func NumberOfTR(doc *goquery.Document) int {
	return doc.Find("tr").Length()
}

// NumberOfHrefs returns the number of href attributes in link elements in the doc
func NumberOfHref(doc *goquery.Document) int {
	count := 0
	doc.Find("link").Each(func(_ int, link *goquery.Selection) {
		if _, exists := link.Attr("href"); exists {
			count++
		}
	})
	return count
}

// NumberOfParagraphs returns the number of paragraph elements in the doc
func NumberOfParagraph(doc *goquery.Document) int {
	return doc.Find("p").Length()
}

// NumberOfScripts returns the number of script elements in the doc
func NumberOfScript(doc *goquery.Document) int {
	return doc.Find("script").Length()
}

// LengthOfTitle returns the length of the title text in the doc
func LengthOfTitle(doc *goquery.Document) int {
	title := doc.Find("title").Text()
	return len(title)
}

// HasH1 checks if the document has an h1 element
func HasH1(doc *goquery.Document) int {
	if doc.Find("h1").Length() > 0 {
		return 1
	}
	return 0
}

// HasH2 checks if the document has an h2 element
func HasH2(doc *goquery.Document) int {
	if doc.Find("h2").Length() > 0 {
		return 1
	}
	return 0
}

// HasH3 checks if the document has an h3 element
func HasH3(doc *goquery.Document) int {
	if doc.Find("h3").Length() > 0 {
		return 1
	}
	return 0
}

// LengthOfText returns the length of text in the document
func LengthOfText(doc *goquery.Document) int {
	return len(doc.Text())
}

// NumberOfClickableButton returns the number of clickable buttons in the document
func NumberOfClickableButton(doc *goquery.Document) int {
	count := 0
	doc.Find("button").Each(func(_ int, button *goquery.Selection) {
		if button.AttrOr("type", "") == "button" {
			count++
		}
	})
	return count
}

// NumberOfA returns the number of <a> tags in the document
func NumberOfA(doc *goquery.Document) int {
	return doc.Find("a").Length()
}

// NumberOfImg returns the number of <img> tags in the document
func NumberOfImg(doc *goquery.Document) int {
	return doc.Find("img").Length()
}

// NumberOfDiv returns the number of <div> tags in the document
func NumberOfDiv(doc *goquery.Document) int {
	return doc.Find("div").Length()
}

// NumberOfFigure returns the number of <figure> tags in the document
func NumberOfFigure(doc *goquery.Document) int {
	return doc.Find("figure").Length()
}

// HasFooter checks if the document has a footer element
func HasFooter(doc *goquery.Document) int {
	if doc.Find("footer").Length() > 0 {
		return 1
	}
	return 0
}

// HasForm checks if the document has a form element
func HasForm(doc *goquery.Document) int {
	if doc.Find("form").Length() > 0 {
		return 1
	}
	return 0
}

// HasTextArea checks if the document has a textarea element
func HasTextArea(doc *goquery.Document) int {
	if doc.Find("textarea").Length() > 0 {
		return 1
	}
	return 0
}

// HasIFrame checks if the document has an iframe element
func HasIFrame(doc *goquery.Document) int {
	if doc.Find("iframe").Length() > 0 {
		return 1
	}
	return 0
}

// HasTextInput checks if the document has a text input element// HasTextInput checks if the document has a text input element
func HasTextInput(doc *goquery.Document) int {
	hasTextInput := 0
	doc.Find("input").Each(func(_ int, input *goquery.Selection) {
		if input.AttrOr("type", "") == "text" {
			hasTextInput = 1
		}
	})
	return hasTextInput
}

// NumberOfMeta returns the number of <meta> tags in the document
func NumberOfMeta(doc *goquery.Document) int {
	return doc.Find("meta").Length()
}

// HasNav checks if the document has a nav element
func HasNav(doc *goquery.Document) int {
	if doc.Find("nav").Length() > 0 {
		return 1
	}
	return 0
}

// HasObject checks if the document has an object element
func HasObject(doc *goquery.Document) int {
	if doc.Find("object").Length() > 0 {
		return 1
	}
	return 0
}

// HasPicture checks if the document has a picture element
func HasPicture(doc *goquery.Document) int {
	if doc.Find("picture").Length() > 0 {
		return 1
	}
	return 0
}

// NumberOfSources returns the number of <source> tags in the document
func NumberOfSources(doc *goquery.Document) int {
	return doc.Find("source").Length()
}

// NumberOfSpan returns the number of <span> tags in the document
func NumberOfSpan(doc *goquery.Document) int {
	return doc.Find("span").Length()
}

// NumberOfTable returns the number of <table> tags in the document
func NumberOfTable(doc *goquery.Document) int {
	return doc.Find("table").Length()
}
