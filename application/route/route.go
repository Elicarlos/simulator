package route
import "errors"

type Route struct {
	ID string 
	ClientId string
	Positions []Position
}

type Position struct {
	Lat float64
	Long float64
}

func(r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New(text: "route id not informad")
	}
	f, err := os.Open(name:"destinations/" + r.ID + ".text")

	if err != nil {
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan(){
		data := string.Split(scanner.Text(), sep=',')
		lat, err := strconv.ParseFloat(data[0], bitSize: 64)
		if err != nil {
			return nil
		}
		long, err := strconv.ParseFloat(data[1], bitSize: 64)
		if err != nil {
			return nil
		}
		r.Positions = append(r.Positions, Position){
			Lat: lat,
			Long: long
		}
	}

	return nil
}