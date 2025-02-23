package parser

import "golang.org/x/net/html"

//isAnchor will validate an Anchor tag
func isAnchor(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "a"
}

//isDiv will validate an Div tag
func isDiv(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "div"
}

//isSpan will validate an span tag
func isSpan(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "span"
}

//isH1 will validate an H1 tag
func isH1(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "h1"
}

//isH2 will validate an H2 tag
func isH2(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "h2"
}

//isH3 will validate an H3 tag
func isH3(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "h3"
}

//isH4 will validate an H4 tag
func isH4(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "h4"
}

//isH5 will validate an H5 tag
func isH5(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "h5"
}

//isH6 will validate an H6 tag
func isH6(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "h6"
}

//isParagraph will validate an Paragraph tag
func isParagraph(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "p"
}

//isImage will validate an Image tag
func isImage(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "img"
}

//isUL will validate an UL tag
func isUL(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "ul"
}

//isOL will validate an OL tag
func isOL(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "ol"
}

//isLI will validate an LI tag
func isLI(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "li"
}

//isScript will validate an Script tag
func isScript(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "script"
}

//isDocTypeOrDocumentType will validate if a tag is DocType or DocumentType
func isDocTypeOrDocumentType(n *html.Node) bool {
	return n.Type == html.DocumentNode || n.Type == html.DoctypeNode
}

//isTextNode will validate if node is TextNode
func isTextNode(n *html.Node) bool {
	if n.Type == html.TextNode && n.DataAtom == 0 {
		return true
	}
	return false
}

//isInput will validate an Image tag
func isInput(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "input"
}

//isVideo will validate an Image tag
func isVideo(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "video"
}

//isVideo will validate an Image tag
func isAudio(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "audio"
}

//isSelect will validate an Image tag
func isSelect(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "select"
}

//isTextArea will validate an Image tag
func isTextArea(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "textarea"
}

//isButton will validate an Image tag
func isButton(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "button"
}

//isIFrame will validate an Image tag
func isIFrame(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "iframe"
}

//isArea will validate an Area tag
func isArea(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "area"
}

//isObject will validate an object tag
func isObject(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "object"
}

//isEmbed will validate an embed tag
func isEmbed(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "embed"
}

//isTrack will validate an embed tag
func isTrack(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "track"
}

//isApplet will validate an embed tag
func isApplet(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "applet"
}
