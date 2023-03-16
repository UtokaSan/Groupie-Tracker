package cmd

type Input struct {
	Input string `json:"input"`
}

type ImageID struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	Genre        string
}
type Test struct {
	Artists  []ImageID   `json:"artists"`
	Location AllLocation `json:"location"`
	Dates    AllDates    `json:"dates"`
}
type ArtistInformation struct {
	Artist   ImageID  `json:"artist"`
	Location Location `json:"location"`
	Dates    Date     `json:"dates"`
}

type AllDateLocation struct {
	Index []struct {
		Name          string
		ID            int                 `json:"id"`
		DatesLocation map[string][]string `json:"datesLocations"`
	} `json:"index"`
}
type AllLocation struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
	} `json:"index"`
}

type AllDates struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}
type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Genre     string
}
type DateLocation struct {
	ID            int                 `json:"id"`
	DatesLocation map[string][]string `json:"datesLocation"`
}

type Tag struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type Toptags struct {
	Tag []Tag `json:"tag"`
}

type AllTags struct {
	Toptags Toptags `json:"toptags"`
}

type Topalbums struct {
	Album []struct {
		Name      string `json:"name"`
		Playcount int    `json:"playcount"`
		Image     []struct {
			Text string `json:"#text"`
			Size string `json:"size"`
		} `json:"image"`
	} `json:"album"`
}
type ListenersArtist struct {
	Stats struct {
		Listeners string `json:"listeners"`
		Playcount string `json:"playcount"`
	} `json:"stats"`
	Bio struct {
		Summary string `json:"summary"`
	} `json:"bio"`
}
type AllInfoArtist struct {
	AllListeners AllListeners `json:"allListeners"`
	AllAlbum     AllAlbum     `json:"allAlbum"`
}
type AllListeners struct {
	ListenersArtist ListenersArtist `json:"artist"`
}
type AllAlbum struct {
	Topalbums Topalbums `json:"topalbums"`
}
