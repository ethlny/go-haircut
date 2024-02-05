package database

import (
	"gohairdresser/structs"
)

var db = SetupDatabase()

// ===== For clients =====
func GetAllClients() ([]structs.Client, error) {
	rows, err := db.Query("SELECT uid, email FROM clients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []structs.Client
	for rows.Next() {
		var c structs.Client
		if err := rows.Scan(&c.UID, &c.Email); err != nil {
			return nil, err
		}
		clients = append(clients, c)
	}
	return clients, nil
}

func GetClientByUID(uid string) (structs.Client, error) {
	var c structs.Client
	err := db.QueryRow("SELECT uid, email FROM clients WHERE uid=?", uid).Scan(&c.UID, &c.Email)
	if err != nil {
		return c, err
	}
	return c, nil
}

func GetClientByEmail(email string) (structs.Client, error) {
	var c structs.Client
	err := db.QueryRow("SELECT uid, email, password FROM clients WHERE email=?", email).Scan(&c.UID, &c.Email, &c.Password)
	if err != nil {
		return c, err
	}
	return c, nil
}

// ===== For hairdressers =====
func GetAllHairdressers() ([]structs.Hairdresser, error) {
	rows, err := db.Query("SELECT uid, salonID, firstName, speciality FROM hairdressers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hairdressers []structs.Hairdresser
	for rows.Next() {
		var h structs.Hairdresser
		if err := rows.Scan(&h.UID, &h.SaloonID, &h.FirstName, &h.Speciality); err != nil {
			return nil, err
		}
		hairdressers = append(hairdressers, h)
	}
	return hairdressers, nil
}

func GetHairdresserByUID(uid string) (structs.Hairdresser, error) {
	var h structs.Hairdresser
	err := db.QueryRow("SELECT uid, salonID, firstName, speciality FROM hairdressers WHERE uid=?", uid).Scan(&h.UID, &h.SaloonID, &h.FirstName, &h.Speciality)
	if err != nil {
		return h, err
	}
	return h, nil
}

// ===== For admins =====
func GetAllAdmins() ([]structs.Admin, error) {
	rows, err := db.Query("SELECT uid, name, email FROM admin")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var admins []structs.Admin
	for rows.Next() {
		var a structs.Admin
		if err := rows.Scan(&a.UID, &a.Name, &a.Email); err != nil {
			return nil, err
		}
		admins = append(admins, a)
	}
	return admins, nil
}

func GetAdminByUID(uid string) (structs.Admin, error) {
	var a structs.Admin
	err := db.QueryRow("SELECT uid, name, email FROM admin WHERE uid=?", uid).Scan(&a.UID, &a.Name, &a.Email)
	if err != nil {
		return a, err
	}
	return a, nil
}

// ===== For hair saloons =====
func GetAllHairSaloons() ([]structs.HairSaloon, error) {
	rows, err := db.Query("SELECT uid, name, address, email, phone, openingTime, closingTime FROM hairSaloon")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var salons []structs.HairSaloon
	for rows.Next() {
		var s structs.HairSaloon
		if err := rows.Scan(&s.UID, &s.Name, &s.Address, &s.Email, &s.Phone, &s.OpeningTime, &s.ClosingTime); err != nil {
			return nil, err
		}
		salons = append(salons, s)
	}
	return salons, nil
}

func GetHairSaloonByUID(uid string) (structs.HairSaloon, error) {
	var s structs.HairSaloon
	err := db.QueryRow("SELECT uid, name, address, email, phone, openingTime, closingTime FROM hairSaloon WHERE uid=?", uid).Scan(&s.UID, &s.Name, &s.Address, &s.Email, &s.Phone, &s.OpeningTime, &s.ClosingTime)
	if err != nil {
		return s, err
	}
	return s, nil
}

func GetSaloonByEmail(email string) (structs.HairSaloon, error) {
	var s structs.HairSaloon
	err := db.QueryRow("SELECT uid, name, address, email, phone, openingTime, closingTime FROM hairSaloon WHERE email=?", email).Scan(&s.UID, &s.Name, &s.Address, &s.Email, &s.Phone, &s.OpeningTime, &s.ClosingTime)
	if err != nil {
		return s, err
	}
	return s, nil
}

// ===== For reservations =====
func GetAllReservations() ([]structs.Reservation, error) {
	rows, err := db.Query("SELECT uid, salonID, clientID, hairdresserID, startHour, endHour, status FROM reservation")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []structs.Reservation
	for rows.Next() {
		var r structs.Reservation
		if err := rows.Scan(&r.UID, &r.SaloonID, &r.ClientID, &r.HairdresserID, &r.StartHour, &r.EndHour, &r.Status); err != nil {
			return nil, err
		}
		reservations = append(reservations, r)
	}
	return reservations, nil
}

func GetReservationByUID(uid string) (structs.Reservation, error) {
	var r structs.Reservation
	err := db.QueryRow("SELECT uid, salonID, clientID, hairdresserID, startHour, endHour, status FROM reservation WHERE uid=?", uid).Scan(&r.UID, &r.SaloonID, &r.ClientID, &r.HairdresserID, &r.StartHour, &r.EndHour, &r.Status)
	if err != nil {
		return r, err
	}
	return r, nil
}

// ===== For schedules =====
func GetAllSchedules() ([]structs.Schedule, error) {
	rows, err := db.Query("SELECT uid, hairdresserID, startHour, endHour, availability FROM schedules")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []structs.Schedule
	for rows.Next() {
		var s structs.Schedule
		if err := rows.Scan(&s.UID, &s.HairdresserID, &s.StartHour, &s.EndHour, &s.Availability); err != nil {
			return nil, err
		}
		schedules = append(schedules, s)
	}
	return schedules, nil
}

func GetScheduleByUID(uid string) (structs.Schedule, error) {
	var s structs.Schedule
	err := db.QueryRow("SELECT uid, hairdresserID, startHour, endHour, availability FROM schedules WHERE uid=?", uid).Scan(&s.UID, &s.HairdresserID, &s.StartHour, &s.EndHour, &s.Availability)
	if err != nil {
		return s, err
	}
	return s, nil
}
