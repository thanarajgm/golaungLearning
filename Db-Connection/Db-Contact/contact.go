package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// 1. Connect to MySQL
	db, err := sql.Open("mysql", "root:November@2025@tcp(localhost:3306)/testdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 2. Contact struct (all columns)
	type Contact struct {
		Id                 string
		AccountId          string
		FirstName          string
		LastName           string
		Salutation         string
		MiddleName         string
		Suffix             string
		Email              string
		Phone              string
		MobilePhone        string
		HomePhone          string
		OtherPhone         string
		Fax                string
		MailingStreet      string
		MailingCity        string
		MailingState       string
		MailingPostalCode  string
		MailingCountry     string
		MailingLatitude    float64
		MailingLongitude   float64
		OtherStreet        string
		OtherCity          string
		OtherState         string
		OtherPostalCode    string
		OtherCountry       string
		OtherLatitude      float64
		OtherLongitude     float64
		Title              string
		Department         string
		Description        string
		CreatedDate        string
		CreatedById        string
		LastModifiedDate   string
		LastModifiedById   string
		IsDeleted          bool
		HasOptedOutOfEmail bool
		DoNotCall          bool
	}

	// 3. Create 10 contact records in a slice
	contacts := []Contact{
		// Already existing first record
		{
			Id:                 "C0001000001",
			AccountId:          "A0001000001",
			FirstName:          "John",
			LastName:           "Doe",
			Salutation:         "Mr.",
			MiddleName:         "M",
			Suffix:             "Jr",
			Email:              "john.doe@example.com",
			Phone:              "9876543210",
			MobilePhone:        "9000001234",
			HomePhone:          "0452-123456",
			OtherPhone:         "0400-111222",
			Fax:                "044-123455",
			MailingStreet:      "123 Main St",
			MailingCity:        "Chennai",
			MailingState:       "TN",
			MailingPostalCode:  "600001",
			MailingCountry:     "India",
			MailingLatitude:    12.9716,
			MailingLongitude:   77.5946,
			OtherStreet:        "456 Second St",
			OtherCity:          "Madurai",
			OtherState:         "TN",
			OtherPostalCode:    "625001",
			OtherCountry:       "India",
			OtherLatitude:      9.9252,
			OtherLongitude:     78.1198,
			Title:              "Manager",
			Department:         "Sales",
			Description:        "Sample contact record",
			CreatedDate:        "2025-01-01 10:00:00",
			CreatedById:        "U0001000001",
			LastModifiedDate:   "2025-01-02 09:00:00",
			LastModifiedById:   "U0001000001",
			IsDeleted:          false,
			HasOptedOutOfEmail: false,
			DoNotCall:          false,
		},
	}

	// 4. Add 9 NEW contacts using loop
	for i := 2; i <= 10; i++ {
		contact := Contact{
			Id:                 fmt.Sprintf("C000100000%d", i),
			AccountId:          fmt.Sprintf("A000100000%d", i),
			FirstName:          fmt.Sprintf("FirstName_%d", i),
			LastName:           fmt.Sprintf("LastName_%d", i),
			Salutation:         "Mr.",
			MiddleName:         "",
			Suffix:             "",
			Email:              fmt.Sprintf("user%d@example.com", i),
			Phone:              fmt.Sprintf("90000000%d", i),
			MobilePhone:        fmt.Sprintf("80000000%d", i),
			HomePhone:          "",
			OtherPhone:         "",
			Fax:                "",
			MailingStreet:      "Sample Street",
			MailingCity:        "Chennai",
			MailingState:       "TN",
			MailingPostalCode:  "600001",
			MailingCountry:     "India",
			MailingLatitude:    12.97,
			MailingLongitude:   77.59,
			OtherStreet:        "",
			OtherCity:          "",
			OtherState:         "",
			OtherPostalCode:    "",
			OtherCountry:       "",
			OtherLatitude:      0,
			OtherLongitude:     0,
			Title:              "Employee",
			Department:         "IT",
			Description:        "Generated contact",
			CreatedDate:        "2025-01-01 10:00:00",
			CreatedById:        "U0001000001",
			LastModifiedDate:   "2025-01-02 09:00:00",
			LastModifiedById:   "U0001000001",
			IsDeleted:          false,
			HasOptedOutOfEmail: false,
			DoNotCall:          false,
		}

		contacts = append(contacts, contact)
	}

	// 5. Insert Query (same as before)
	insertQuery := `
        INSERT INTO Contact (
            Id, AccountId, FirstName, LastName, Salutation, MiddleName, Suffix,
            Email, Phone, MobilePhone, HomePhone, OtherPhone, Fax,
            MailingStreet, MailingCity, MailingState, MailingPostalCode,
            MailingCountry, MailingLatitude, MailingLongitude,
            OtherStreet, OtherCity, OtherState, OtherPostalCode, OtherCountry,
            OtherLatitude, OtherLongitude,
            Title, Department, Description,
            CreatedDate, CreatedById, LastModifiedDate, LastModifiedById,
            IsDeleted, HasOptedOutOfEmail, DoNotCall
        )
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 
                ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `

	// 6. Insert all contacts using loop
	for _, c := range contacts {
		_, err := db.Exec(insertQuery,
			c.Id, c.AccountId, c.FirstName, c.LastName, c.Salutation,
			c.MiddleName, c.Suffix, c.Email, c.Phone, c.MobilePhone,
			c.HomePhone, c.OtherPhone, c.Fax, c.MailingStreet, c.MailingCity,
			c.MailingState, c.MailingPostalCode, c.MailingCountry,
			c.MailingLatitude, c.MailingLongitude, c.OtherStreet, c.OtherCity,
			c.OtherState, c.OtherPostalCode, c.OtherCountry, c.OtherLatitude,
			c.OtherLongitude, c.Title, c.Department, c.Description,
			c.CreatedDate, c.CreatedById, c.LastModifiedDate, c.LastModifiedById,
			c.IsDeleted, c.HasOptedOutOfEmail, c.DoNotCall,
		)

		if err != nil {
			fmt.Println("Insert error:", err)
			continue
		}

		fmt.Println("Inserted:", c.Id)
	}

	fmt.Println("All 10 contacts inserted successfully!")
}
