package analyzerapp

//Divtag is the Div based rule results
type Divtag struct {
	Div         string
	Wc111Aria6  string
	Wc111Aria10 string
}

//Anchortag is the anchor based rule results
type Anchortag struct {
	Anchor      string
	Wc111Aria6  string
	Wc111Aria10 string
}

//Audiotag is the audio based rule results
type Audiotag struct {
	Audio       string
	Wc111Aria6  string
	Wc111Aria10 string
	Wc121H96    string
	Wc121ARIA   string
}

//Buttontag is the button based rule results
type Buttontag struct {
	Button      string
	Wc111Aria6  string
	Wc111Aria10 string
}

//Imgtag is the img based rule results
type Imgtag struct {
	Img         string
	Wc111Aria6  string
	Wc111Aria10 string
	Wc111G94    string
	Wc111H2     string
}

//Inputtag is the input based rule results
type Inputtag struct {
	Input       string
	Wc111Aria6  string
	Wc111Aria10 string
}

//Selecttag is the select based rule results
type Selecttag struct {
	Select      string
	Wc111Aria6  string
	Wc111Aria10 string
}

//Textareatag is the textarea based rule results
type Textareatag struct {
	Textarea    string
	Wc111Aria6  string
	Wc111Aria10 string
}

//Videotag is the video based rule results
type Videotag struct {
	Video       string
	Wc111Aria6  string
	Wc111Aria10 string
}

//Paratag is the para based rule results
type Paratag struct {
	Para        string
	Wc111Aria6  string
	Wc111Aria10 string
}

//Iframetag is the iframe based rule results
type Iframetag struct {
	Iframe      string
	Wc111Aria6  string
	Wc111Aria10 string
}

//Areatag is the area based rule results
type Areatag struct {
	Area     string
	Wc111G94 string
	Wc111H2  string
}

//Objecttag is the object based rule results
type Objecttag struct {
	Object    string
	Wc121H96  string
	Wc121ARIA string
	Wc111H53  string
	Wc121G166 string
}

//Embedtag is the embed based rule results
type Embedtag struct {
	Embed     string
	Wc121H96  string
	Wc121ARIA string
	Wc111H53  string
	Wc121G166 string
}

//Tracktag is the track based rule results
type Tracktag struct {
	Track      string
	Wc121H96   string
	Wc121ARIA  string
	Wc111G94   string
	Wc111H2    string
	Wcag122G87 string
}

//Applettag is the track based rule results
type Applettag struct {
	Applet    string
	Wc121H96  string
	Wc121ARIA string
	Wc111G94  string
	Wc111H2   string
	Wc111H35  string
}
