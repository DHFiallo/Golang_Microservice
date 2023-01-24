package data

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"gioui.org/layout"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var FILE_NAME = "data/UserInformation[22].csv"

// User struct defining structure for API
type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first"`
	LastName    string `json:"last"`
	Email       string `json:"email"`
	Profession  string
	DateCreated string
	Country     string
	City        string
}

type (
	C = layout.Context
	D = layout.Dimensions
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Heading text for UI
var headingText = []string{"ID", "FirstName", "LastName", "Email", "Profession", "DateCreated", "Country", "City"}

// Prints everything out
func Start() {
	records, err, f := readData(FILE_NAME)

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		user := User{
			ID:          id,
			FirstName:   record[1],
			LastName:    record[2],
			Email:       record[3],
			Profession:  record[4],
			DateCreated: record[5],
			Country:     record[6],
			City:        record[7],
		}
		fmt.Printf("%d %s %s %s %s %s %s %s\n", user.ID, user.FirstName, user.LastName, user.Email, user.Profession, user.DateCreated, user.Country, user.City)
	}

	defer f.Close()
}

// Prints out profession
func GetProfession(job string) []User {

	records, err, f := readData(FILE_NAME)

	if err != nil {
		log.Fatal(err)
	}

	userList := []User{}
	for _, record := range records {
		if record[4] == job {
			id, _ := strconv.Atoi(record[0])
			user := User{
				ID:          id,
				FirstName:   record[1],
				LastName:    record[2],
				Email:       record[3],
				Profession:  record[4],
				DateCreated: record[5],
				Country:     record[6],
				City:        record[7],
			}
			userList = append(userList, user)
			//fmt.Printf("This is an ID %d", id)

			//fmt.Printf("%d %s %s %s %s %s %s %s\n", user.ID, user.FirstName, user.LastName, user.Email, user.Profession, user.DateCreated, user.Country, user.City)
		}
	}
	//j for job
	startWindow(userList, 'j')
	defer f.Close()
	return userList
}

// Prints out every user between date range
func GetUsersBetweenDates(date1 time.Time, date2 time.Time) []User {
	records, err, f := readData(FILE_NAME)

	layout := "2006-01-02" //YYYY-MM-DD

	//date2 is latest time (not validated so assuming best case scenario)

	if err != nil {
		log.Fatal(err)
	}

	userList := []User{}
	for _, record := range records {

		recordDate, _ := time.Parse(layout, record[5])
		if recordDate.Before(date2) && date1.Before(recordDate) {

			id, _ := strconv.Atoi(record[0])
			user := User{
				ID:          id,
				FirstName:   record[1],
				LastName:    record[2],
				Email:       record[3],
				Profession:  record[4],
				DateCreated: record[5],
				Country:     record[6],
				City:        record[7],
			}
			userList = append(userList, user)
			//fmt.Printf("%d %s %s %s %s %s %s %s\n", user.ID, user.FirstName, user.LastName, user.Email, user.Profession, user.DateCreated, user.Country, user.City)
		}
	}
	//d for date
	startWindow(userList, 'd')
	defer f.Close()
	return userList
}

func GetSpecificPerson(first string, last string) []User {
	records, err, f := readData(FILE_NAME)

	if err != nil {
		log.Fatal(err)
	}

	userList := []User{}
	for _, record := range records {
		//Use EqualFold as it can compares case insensitively. Didn't use toLower or toUpper as it can have issues
		if strings.EqualFold(first, record[1]) && strings.EqualFold(last, record[2]) {

			id, _ := strconv.Atoi(record[0])
			user := User{
				ID:          id,
				FirstName:   record[1],
				LastName:    record[2],
				Email:       record[3],
				Profession:  record[4],
				DateCreated: record[5],
				Country:     record[6],
				City:        record[7],
			}
			userList = append(userList, user)
			//fmt.Printf("%d %s %s %s %s %s %s %s\n", user.ID, user.FirstName, user.LastName, user.Email, user.Profession, user.DateCreated, user.Country, user.City)
		}
	}
	//p for person
	startWindow(userList, 'p')
	defer f.Close()
	return userList
}

func (u *User) UpdateUser(id int) {
	records, err, f := readData(FILE_NAME)
	defer f.Close()

	csvFile, err := os.Create(FILE_NAME)
	w := csv.NewWriter(csvFile)

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		recordID, _ := strconv.Atoi(record[0])
		if recordID == id {
			record[1] = u.FirstName
			record[2] = u.LastName
			record[3] = u.Email
			record[4] = u.Profession
			record[5] = u.DateCreated
			record[6] = u.Country
			record[7] = u.City
			break
		}
	}

	w.WriteAll(records)

	//u for update
	GetSpecificPerson(u.FirstName, u.LastName)
	defer f.Close()

	defer w.Flush()
}

func readData(fileName string) ([][]string, error, *os.File) {

	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err, f
	}

	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err, f
	}

	return records, nil, f
}

func (u *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

// this was made for some abstraction
type Users []*User

func (u *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

var ErrUserNotFound = fmt.Errorf("User not found")

func startWindow(userList []User, r rune) {
	var app = tview.NewApplication()

	list := tview.NewList()
	for _, user := range userList {
		text := user.FirstName + " " + user.LastName + " | " + user.Email + " | " + user.Profession + " | " + user.DateCreated + " | " + user.City + "," + user.Country

		list.AddItem(strconv.Itoa(user.ID), text, r, nil)
		//AddItem("List item 1", "Some explanatory text", 'a', nil).
	}
	list.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
	})

	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			app.Stop()
		}
		return event
	})
}

// adds user to array (change so it changes csv)
/*
func AddUser(u *User) {
	u.ID = getNextID()
	userList = append(userList, u)
}
*/

/*
func UpdateUser(id int, u *User) error {
	_, pos, err := findUser(id)

	if err != nil {
		return err
	}

	u.ID = id
	userList[pos] = u

	return err
}
*/

//Tried to use Gio here to make a gui to show data
/*
func startWindow(userList []User) {
	go func() {
		//creates window

		w := app.NewWindow(
			app.Title("User Information"),
			app.Size(unit.Dp(650), unit.Dp(600)),
		)

		if err := draw(w, userList); err != nil {
			log.Fatal(err)
		}

		/*if err := loop(w, userList); err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}()
	app.Main()
}

// The main draw function
func draw(w *app.Window, userList []User) error {
	// y-position for text
	var scrollY unit.Dp = 0

	// width of text area
	var textWidth unit.Dp = 550

	// fontSize
	var fontSize unit.Sp = 35

	// Are we auto scrolling?
	var autoscroll bool = false
	var autospeed unit.Dp = 1

	// th defines the material design style
	th := material.NewTheme(gofont.Collection())

	// ops are the operations from the UI
	var ops op.Ops

	// Listen for events from the window.
	for windowEvent := range w.Events() {
		switch winE := windowEvent.(type) {

		// Should we draw a new frame?
		case system.FrameEvent:

			// ---------- Handle input ----------
			// Time to deal with inputs since last frame.
			// Since we use one global eventArea, with Tag: 0
			// we here call gtx.Events(0) to get these events.

			gtx := layout.NewContext(&ops, winE)
			for _, gtxEvent := range gtx.Events(0) {

				// To set how large each change is
				var stepSize unit.Dp = 1

				switch gtxE := gtxEvent.(type) {

				// Any key
				case key.EditEvent:
					// To increase the fontsize
					if gtxE.Text == "+" {
						fontSize = fontSize + unit.Sp(stepSize)
					}
					// To decrease the fontsize
					if gtxE.Text == "-" {
						fontSize = fontSize - unit.Sp(stepSize)
					}

				// Only specified keys, defined in key.InputOp below
				case key.Event:
					// For better control, we only care about pressing the key down, not releasing it up
					if gtxE.State == key.Press {
						// Start/Stop
						if gtxE.Name == "Space" {
							autoscroll = !autoscroll
							if autospeed == 0 {
								autoscroll = true
								autospeed++
							}
						}
					}

				// A mouse event?
				case pointer.Event:
					// Are we scrolling?
					if gtxE.Type == pointer.Scroll {
						if gtxE.Modifiers == key.ModShift {
							stepSize = 3
						}
						// Increment scrollY with gtxE.Scroll.Y
						scrollY = scrollY + unit.Dp(gtxE.Scroll.Y)*stepSize
						if scrollY < 0 {
							scrollY = 0
						}
					}
				}
			}

			// ---------- LAYOUT ----------
			// First we layout the user interface.
			// Afterwards we add an eventArea.
			// Let's start with a background color
			paint.Fill(&ops, color.NRGBA{R: 0xff, G: 0xfe, B: 0xe0, A: 0xff})

			// ---------- THE SCROLLING TEXT ----------
			// First, check if we should autoscroll
			// That's done by increasing the value of scrollY
			if autoscroll {
				scrollY = scrollY + autospeed
				op.InvalidateOp{At: gtx.Now.Add(time.Second * 2 / 100)}.Add(&ops)
			}
			// Then we use scrollY to control the distance from the top of the screen to the first element.
			// We visualize the text using a list where each paragraph is a separate item.
			var visList = layout.List{
				Axis: layout.Vertical,
				Position: layout.Position{
					Offset: int(scrollY),
				},
			}

			// ---------- MARGINS ----------
			// Margins
			var marginWidth unit.Dp
			marginWidth = (unit.Dp(gtx.Constraints.Max.X) - textWidth) / 2
			margins := layout.Inset{
				Left:   marginWidth,
				Right:  marginWidth,
				Top:    unit.Dp(0),
				Bottom: unit.Dp(0),
			}

			// ---------- LIST WITHIN MARGINS ----------
			// 1) First the margins ...

			//v := reflect.ValueOf(userList[0])

			margins.Layout(gtx,
				func(gtx C) D {
					// 2) ... then the list inside those margins ...
					return visList.Layout(gtx, 100,
						// 3) ... where each paragraph is a separate item
						func(gtx C, index int) D {
							// One label per paragraph
							str := fmt.Sprintf("%v", userList)
							line := material.Label(th, unit.Sp(float32(fontSize)), str)
							// The text is centered
							line.Alignment = text.Middle
							// Return the laid out paragraph
							return line.Layout(gtx)
						},
					)
				},
			)

			// ---------- COLLECT INPUT ----------
			// Create an eventArea to collect events. It has the same size as the full windodw.
			// First we Push() it on the stack, then add code to catch keys and pointers
			// Finally we Pop() it
			eventArea := clip.Rect(
				image.Rectangle{
					// From top left
					Min: image.Point{0, 0},
					// To bottom right
					Max: image.Point{gtx.Constraints.Max.X, gtx.Constraints.Max.Y},
				},
			).Push(gtx.Ops)

			// Since Gio is stateless we must Tag events, to make sure we know where they came from.
			// Such a tag can anything really, so we simply use Tag: 0.
			// Later we retireve these events with gtx.Events(0)

			// 1) We first add a pointer.InputOp to catch scrolling:
			pointer.InputOp{
				Types: pointer.Scroll,
				Tag:   0, // Use Tag: 0 as the event routing tag, and retireve it through gtx.Events(0)
				// ScrollBounds sets bounds on scrolling, and we want it to be non-zero.
				// In practice it seldom reached 100, so [MinInt8,MaxInt8] or [-128,127] should be enough
				ScrollBounds: image.Rectangle{
					Min: image.Point{
						X: 0,
						Y: math.MinInt8, //-128
					},
					Max: image.Point{
						X: 0,
						Y: math.MaxInt8, //+127
					},
				},
			}.Add(gtx.Ops)

			// 2) Next we add key.FocusOp,
			// Needed for general keybaord input, except the ones defined explicitly in key.InputOp
			// These inputs are retrieved as key.EditEvent
			key.FocusOp{
				Tag: 0, // Use Tag: 0 as the event routing tag, and retireve it through gtx.Events(0)
			}.Add(gtx.Ops)

			// 3) Finally we add key.InputOp to catch specific keys
			// (Shift) means an optional Shift
			// These inputs are retrieved as key.Event
			key.InputOp{
				Keys: key.Set("(Shift)-F|(Shift)-S|(Shift)-U|(Shift)-D|(Shift)-J|(Shift)-K|(Shift)-W|(Shift)-N|Space"),
				Tag:  0, // Use Tag: 0 as the event routing tag, and retireve it through gtx.Events(0)
			}.Add(gtx.Ops)

			// Finally Pop() the eventArea from the stack
			eventArea.Pop()

			// ---------- FINALIZE ----------
			// Frame completes the FrameEvent by drawing the graphical operations from ops into the window.
			winE.Frame(&ops)

		// Should we shut down?
		case system.DestroyEvent:
			return winE.Err

		}

	}
	return nil
}

func loop(w *app.Window, userList []User) error {
	th := material.NewTheme(gofont.Collection())
	var (
		ops  op.Ops
		grid component.GridState
	)
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			op.InvalidateOp{}.Add(gtx.Ops)
			layoutTable(th, gtx, userList, &grid)
			e.Frame(gtx.Ops)
		}
	}
}

func layoutTable(th *material.Theme, gtx C, userList []User, grid *component.GridState) D {
	// Configure width based on available space and a minimum size.
	minSize := gtx.Dp(unit.Dp(200))
	border := widget.Border{
		Color: color.NRGBA{A: 255},
		Width: unit.Dp(1),
	}

	inset := layout.UniformInset(unit.Dp(2))

	// Configure a label styled to be a heading.
	headingLabel := material.Body1(th, "")
	headingLabel.Font.Weight = text.Bold
	headingLabel.Alignment = text.Middle
	headingLabel.MaxLines = 1

	// Configure a label styled to be a data element.
	dataLabel := material.Body1(th, "")
	dataLabel.Font.Variant = "Mono"
	dataLabel.MaxLines = 1
	dataLabel.Alignment = text.End

	// Measure the height of a heading row.
	orig := gtx.Constraints
	gtx.Constraints.Min = image.Point{}
	macro := op.Record(gtx.Ops)
	dims := inset.Layout(gtx, headingLabel.Layout)
	_ = macro.Stop()
	gtx.Constraints = orig

	return component.Table(th, grid).Layout(gtx, len(userList), 4,
		func(axis layout.Axis, index, constraint int) int {
			widthUnit := max(int(float32(constraint)/3), minSize)
			switch axis {
			case layout.Horizontal:
				switch index {
				case 0, 1:
					return int(widthUnit)
				case 2, 3:
					return int(widthUnit / 2)
				default:
					return 0
				}
			default:
				return dims.Size.Y
			}
		},
		func(gtx C, col int) D {
			return border.Layout(gtx, func(gtx C) D {
				return inset.Layout(gtx, func(gtx C) D {
					headingLabel.Text = headingText[col]
					return headingLabel.Layout(gtx)
				})
			})
		},
		func(gtx C, row, col int) D {
			return inset.Layout(gtx, func(gtx C) D {
				switch col {
				case 0:
					dataLabel.Text = ""
				case 1:
					dataLabel.Text = ""
				case 2:
					dataLabel.Text = ""
				case 3:
					dataLabel.Text = ""
				case 4:
					dataLabel.Text = ""
				case 5:
					dataLabel.Text = ""
				case 6:
					dataLabel.Text = ""
				case 7:
					dataLabel.Text = ""
				case 8:
					dataLabel.Text = ""

				}
				return dataLabel.Layout(gtx)
			})
		},
	)
}
*/
