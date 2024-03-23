package api

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const system_option_count = 4

func (a *AppServer) HomeHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) *ApiError {
	return nil
}

func (a *AppServer) WelcomeHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) *ApiError {
	// sessionId := r.PostForm.Get("sessionId")
	// serviceCode := r.PostForm.Get("serviceCode")

	response := "CON "

	err := r.ParseForm()
	if err != nil {
		response = "System error"
		return nil
	}

	phoneNumber := r.Form.Get("phoneNumber")
	text := r.Form.Get("text")

parent:
	switch text {
	case "":
		// present auth page
		response += "Select an option\n"
		response += "1. Enter Access PIN\n"
		response += "2. Initailize Account\n"
		response += "3. Cancel"
	case "1":
		response += "Enter Access Code\n"
		response += "1. Cancel"
	case "2":
		companies, err := a.db.ListCompanies(ctx)
		if err != nil {
			fmt.Printf("\n\n\n\n")
			fmt.Println(err)
			fmt.Printf("\n\n\n\n")
			response += "System error"
			break
		}
		response += "Select Company\n"

		for i, v := range companies {
			response += fmt.Sprintf("%d. %s\n", i+1, v.Name)
		}
		response += fmt.Sprintf("%d. Cancel", len(companies)+1)
	case "3":
		// cancel session
	default:

		text = strings.ReplaceAll(text, `"`, "")
		// selecting company
		init_level_1, _ := regexp.MatchString(`^2\*\d$`, text)
		if init_level_1 {
			// get company packages
			parts := strings.Split(text, "*")
			id, err := strconv.Atoi(parts[len(parts)-1])
			if err != nil {
				response += "System error"
				break
			}

			response += "Select Package\n"
			packages, err := a.db.GetCompanyPackages(ctx, int32(id))
			if err != nil {
				response += "System error"
				break
			}

			for i, v := range packages {
				response += fmt.Sprintf("%d. %s (%s)\n", i+1, v.Name, v.Price.String)
			}
			response += fmt.Sprintf("%d. Cancel", len(packages)+1)
			break
		}

		// handle login pin entering
		authed_level_1, _ := regexp.MatchString(`^1\*\d{4}$`, text)
		if authed_level_1 {
			// get company packages
			parts := strings.Split(text, "*")
			pin, err := strconv.Atoi(parts[len(parts)-1])
			if err != nil {
				response += "System error"
				break
			}

			// get user
			holder, err := a.db.GetUser(ctx, phoneNumber)
			if err != nil {
				response += "User not found"
				break
			}
			if holder.Code != int32(pin) {
				response += "Invalid pin"
				break
			}

			// valid
			// present the user with services menu
			response += "Select a service\n"
			response += "1. Account Details\n"
			response += "2. Credit Account\n"
			response += "3. Initialize Claim\n"

			// add company custom services

			// get holder package services.
			services, service_fetch_err := a.db.GetUserCompantServices(ctx, holder.ID)
			if service_fetch_err != nil {
				response += "System error"
				break
			}

			next := system_option_count
			for i, v := range services {
				next += i
				response += fmt.Sprintf("%d: %s\n", next, v.Name)
			}
			response += fmt.Sprintf("%d. Cancel", next+1)
		}

		authed_level_2, _ := regexp.MatchString(`^1\*\d{4}\*\d$`, text)
		if authed_level_2 {
			// get select service
			parts := strings.Split(text, "*")
			value, err := strconv.Atoi(parts[len(parts)-1])
			if err != nil {
				response += "System error"
				break
			}

			switch value {
			case 1:
				// make request to node server
				response += "Select an option\n"
				response += "1. View Details\n"
				response += "2. Check Credit\n"
				response += "3. View Teams\n"
				response += "4. Deactivate\n"
				response += "5. Cancel"
				break parent
			case 2:
				// credit account
				response += "Transfer Money from\n"
				response += "1. Registered Mobile wallet\n"
				response += "2. Other Mobile wallet\n"
				response += "3. Bank\n"
				response += "4. Cancel"
				break parent
			case 3:
				response += "Enter reason for claim"
				break parent
			}

		}

		account_details_level, _ := regexp.MatchString(`^1\*\d{4}\*1\*\d$`, text)
		if account_details_level {
			parts := strings.Split(text, "*")
			value, err := strconv.Atoi(parts[len(parts)-1])
			if err != nil {
				response += "System error"
				break
			}
			// get user data
			holder, err := a.db.GetUser(ctx, phoneNumber)
			if err != nil {
				response += "User not found"
				break
			}

			switch value {
			case 1:
				response += fmt.Sprintf("Name: %s\n", holder.Name)
				response += fmt.Sprintf("Phone Number: %s\n", holder.Telnumber)
				response += fmt.Sprintf("Access Code: %d\n", holder.Code)
			case 2:
			case 3:
			}
		}

		claim_level_1, _ := regexp.MatchString(`^1\*\d{4}\*3.*`, text)
		if claim_level_1 {
			parts := strings.Split(text, "*")
			reason := parts[len(parts)-1]

			// pass reason to company for outreach

			response = reason
			break
		}

		credit_account_level_1, _ := regexp.MatchString(`^1\*\d{4}\*2\*\d$`, text)
		if credit_account_level_1 {
			parts := strings.Split(text, "*")
			value, err := strconv.Atoi(parts[len(parts)-1])
			if err != nil {
				response += "System error"
				break
			}
			switch value {
			case 1:
				// get user data
				holder, err := a.db.GetUser(ctx, phoneNumber)
				if err != nil {
					response += "User not found"
					break
				}

				response += fmt.Sprintf("Funding account using %s\n", holder.Telnumber)
				response += "1. Confirm\n"
				response += "2. cancel"
				break parent
			case 2:
				response += "Enter Wallet Number"
				break parent
			case 3:
				response += "Select Bank\n"
				response += "1. Stanbic\n"
				response += "1. Standard Chartered\n"
				break parent
			case 4:
				// cancel
			}
		}

		// from registered wallet
		credit_account_level_2, _ := regexp.MatchString(`^1\*\d{4}\*2\*1\*\d$`, text)
		if credit_account_level_2 {
			parts := strings.Split(text, "*")
			value, err := strconv.Atoi(parts[len(parts)-1])
			if err != nil {
				response += "System error"
				break
			}
			switch value {
			case 1:
				// make transaction
			case 2:
				// cancel
			}
		}

		// handling custom company services request
		custom_services, _ := regexp.MatchString(`^1\*\d{4}\*\d$`, text)
		if custom_services {
			parts := strings.Split(text, "*")
			id, err := strconv.Atoi(parts[len(parts)-1])
			if err != nil {
				response += "System error"
				break
			}

			// get user
			holder, err := a.db.GetUser(ctx, phoneNumber)
			if err != nil {
				response += "User not found"
				break
			}

			// get services
			services, fetch_err := a.db.GetUserCompantServices(ctx, holder.ID)
			if fetch_err != nil {
				response += "System error"
				break
			}

			service_id := id - system_option_count
			selected_service := services[service_id]

			response += fmt.Sprintf("%s\n", selected_service.Name)
			response += fmt.Sprintf("%s\n", selected_service.Description.String)
			break
		}

	}

	// check if user have subscription
	// holder, fetch_err := a.db.GetUser(ctx, phoneNumber)
	// if !errors.Is(fetch_err, sql.ErrNoRows) {
	// 	w.Header().Add("Content-Type", "text/plain")
	// 	w.Write([]byte("System error"))
	// 	return nil
	// }

	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(response))
	return nil
}

// func (a *AppServer) AuthHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) *ApiError {
// 	phoneNumber := r.PostForm.Get("phoneNumber")
// 	text := r.PostForm.Get("text")
// 	code, err := strconv.Atoi(text)
// 	if err != nil {
// 		w.Header().Add("Content-Type", "text/plain")
// 		w.Write([]byte("Invalid Input"))
// 	}

// 	holder, f_err := a.db.GetUser(ctx, phoneNumber)
// 	if f_err != nil {
// 		w.Header().Add("Content-Type", "text/plain")
// 		w.Write([]byte("Invalid Input"))
// 	}

// 	if holder.Code != int32(code) {
// 		w.Header().Add("Content-Type", "text/plain")
// 		w.Write([]byte("Invalid Input"))
// 	}

// 	// get holder package services.
// 	services, service_fetch_err := a.db.GetUserCompantServices(ctx, holder.ID)
// 	if service_fetch_err != nil {
// 		w.Header().Add("Content-Type", "text/plain")
// 		w.Write([]byte("System Error"))
// 	}

// 	response := "Select an option\n"
// 	response += "1. Account Details\n"
// 	response += "2. Credit Account\n"
// 	response += "3. Initialize Claim\n"

// 	next := 4
// 	for i, v := range services {
// 		response += fmt.Sprintf("%d: %s\n", next, v.Name)

// 	}

// 	return nil
// }

// func (a *AppServer) SubcriptionInitializationHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) *ApiError {
// 	// phoneNumber := r.PostForm.Get("phoneNumber")
// 	text := r.PostForm.Get("text")
// 	val, err := strconv.Atoi(text)
// 	if err != nil {
// 		w.Write([]byte("Invalid Input"))
// 	}
// 	switch val {
// 	case 1:
// 		// get code
// 	case 2:
// 		// cancellation
// 	}
// 	return nil
// }
