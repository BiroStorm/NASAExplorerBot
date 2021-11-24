package main

/*APOD is a struct for the NASA Astronomy Picture Of the Day
should be used for https://api.nasa.gov/planetary/apod
*/
type APOD struct {
	Copyright string `json:"copyright,omitempty"`
	Date      string `json:"date"`
	Text      string `json:"explanation"`
	URLHD     string `json:"hdurl"`
	Type      string `json:"media_type"`
	Title     string `json:"title"`
	URL       string `json:"url"`
}
