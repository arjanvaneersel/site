package meetup

import "time"

// Event is the data model for meetup events as specified in the v3 API description
// https://www.meetup.com/meetup_api/docs/:urlname/events/#list
type Event struct {
	AttendanceCount   int           `json:"attendance_count"`
	AttendanceSample  string        `json:"attendance_sample"`
	CommentCount      int           `json:"comment_count"`
	Created           MeetupTime    `json:"created"`
	Description       string        `json:"description"`
	DescriptionImages string        `json:"description_images"`
	Duration          time.Duration `json:"duration"`
	EventHosts        []struct {
		ID    int    `json:"id"`
		Into  string `json:"intro"`
		Name  string `json:"name"`
		Photo struct {
			BaseURL     string `json:"base_url"`
			HighresLink string `json:"highres_link"`
			ID          int    `json:"id"`
			PhotoLink   string `json:"photo_link"`
			ThumbLink   string `json:"thumb_link"`
			Type        string `json:"type"`
		} `json:"photo"`
	} `json:"event_hosts"`
	Featured      bool `json:"featured"`
	FeaturedPhoto struct {
		BaseURL     string `json:"base_url"`
		HighresLink string `json:"highres_link"`
		ID          int    `json:"id"`
		PhotoLink   string `json:"photo_link"`
		ThumbLink   string `json:"thumb_link"`
		Type        string `json:"type"`
	} `json:"featured_photo"`
	Fee struct {
		Accepts     string  `json:"accepts"`
		Amount      float64 `json:"amount"`
		Currency    string  `json:"currency"`
		Description string  `json:"per-person"`
		Label       string  `json:"label"`
		Required    bool    `json:"required"`
	} `json:"fee"`
	FeeOptions struct {
		Currencies []struct {
			Code    string `json:"code"`
			Default bool   `json:"default"`
		} `json:"currencies"`
		IsSetup   bool   `json:"is_setup"`
		SetupLink string `json:"setup_link"`
		Type      string `json:"type"`
	} `json:"fee_options"`
	Group struct {
		Category struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			ShortName string `json:"shortname"`
			SortName  string `json:"sort_name"`
		} `json:"category"`
		ID       int `json:"id"`
		JoinInfo struct {
			PhotoReq     bool     `json:"photo_req"`
			Questions    []string `json:"questions"`
			QuestionsReq bool     `json:"questions_req"`
		} `json:"join_info"`
		JoinMode string     `json:"join_mode"`
		Created  MeetupTime `json:"created"`
		KeyPhoto struct {
			BaseURL     string `json:"base_url"`
			HighresLink string `json:"highres_link"`
			ID          int    `json:"id"`
			PhotoLink   string `json:"photo_link"`
			ThumbLink   string `json:"thumb_link"`
			Type        string `json:"type"`
		} `json:"key_photo"`
		Lat               float64 `json:"lat"`
		LocalizedLocation string  `json:"localized_location"`
		Lon               float64 `json:"lon"`
		MembershipDues    struct {
			Currency            string   `json:"currency"`
			Fee                 int64    `json:"fee"`
			FeeDesc             string   `json:"fee_desc"`
			Methods             []string `json:"methods"`
			Reasons             []string `json:"reasons"`
			ReasonsOther        []string `json:"reasons_other"`
			RefundPolicy        string   `json:"refund_policy"`
			Required            bool     `json:"required"`
			RequiredTo          string   `json:"required_to"`
			SelfPaymentRequired bool     `json:"self_payment_required"`
			TrialDays           int      `json:"trial_days"`
		} `json:"membership_dues"`
		MetaCategory struct {
			BestTopics  []string `json:"best_topics"`
			CategoryIDs []int    `json:"category_ids"`
			ID          int      `json:"id"`
			Name        string   `json:"name"`
			Photo       string   `json:"photo"`
			Shortname   string   `json:"shortname"`
			SortName    string   `json:"sort_name"`
		} `json:"meta_category"`
		Name           string `json:"name"`
		PastEventCount int    `json:"past_event_count"`
		Photo          struct {
			BaseURL     string `json:"base_url"`
			HighresLink string `json:"highres_link"`
			ID          int    `json:"id"`
			PhotoLink   string `json:"photo_link"`
			ThumbLink   string `json:"thumb_link"`
			Type        string `json:"type"`
		} `json:"photo"`
		PhotoGradient struct {
			CompositeColor int `json:"composite_color"`
			DarkColor      int `json:"dark_color"`
			ID             int `json:"id"`
			LightColor     int `json:"light_color"`
		} `json:"photo_gradient"`
		Region string `json:"en_US"`
		Self   struct {
			Actions        []string `json:"actions"`
			MembershipDues struct {
				Currency            string   `json:"currency"`
				Fee                 int64    `json:"fee"`
				FeeDesc             string   `json:"fee_desc"`
				Methods             []string `json:"methods"`
				Reasons             []string `json:"reasons"`
				ReasonsOther        []string `json:"reasons_other"`
				RefundPolicy        string   `json:"refund_policy"`
				Required            bool     `json:"required"`
				RequiredTo          string   `json:"required_to"`
				SelfPaymentRequired bool     `json:"self_payment_required"`
				TrialDays           int      `json:"trial_days"`
			} `json:"membership_dues"`
			Profile string `json:"profile"`
			Status  string `json:"status"`
		} `json:"self"`
		Topic []struct {
			ID     int    `json:"id"`
			Lang   string `json:"lang"`
			Name   string `json:"name"`
			URLKey string `json:"urlkey"`
		} `json:"topics"`
		URLName    string `json:"urlname"`
		Visibility string `json:"visibility"`
		Who        string `json:"who"`
	} `json:"group"`
	HowToFindUS             string `json:"how_to_find_us"`
	ID                      string `json:"id"`
	Link                    string `json:"link"`
	LocalDate               string `json:"local_date"`
	LocalTime               string `json:"local_time"`
	ManualAttendanceCount   int    `json:"manual_attendance_count"`
	Name                    string `json:"name"`
	PastEventCountInclusive int    `json:"past_event_count_inclusive"`
	PhotoAlbum              struct {
		Event struct {
			ID            int        `json:"id"`
			Name          string     `json:"name"`
			NoRSVPCount   bool       `json:"no_rsvp_count"`
			Time          MeetupTime `json:"time"`
			UTCOffset     int64      `json:"utc_offset"`
			WaitlistCount int        `json:"waitlist_count"`
			YesRSVPCount  int        `json:"yes_rsvp_count"`
		} `json:"event"`
		ID          int `json:"id"`
		PhotoCount  int `json:"photo_count"`
		PhotoSample struct {
			BaseURL     string `json:"base_url"`
			HighresLink string `json:"highres_link"`
			ID          int    `json:"id"`
			PhotoLink   string `json:"photo_link"`
			ThumbLink   string `json:"thumb_link"`
			Type        string `json:"type"`
		} `json:"photo_sample"`
		Title string `json:"title"`
	} `json:"photo_album"`
	PlainTextDescription       string    `json:"plain_text_description"`
	PlanTextNoImagesDecription string    `json:"plain_text_no_images_description"`
	RSVPCloseOffset            time.Time `json:"rsvp_close_offset"`
	RSVPLimit                  int       `json:"rsvp_limit"`
	RSVPOpenOffset             time.Time `json:"rsvp_open_offset"`
	RSVPRules                  struct {
		CloseTime    time.Time `json:"close_time"`
		Closed       bool      `json:"closed"`
		GuestLimit   int       `json:"guest_limit"`
		OpenTime     time.Time `json:"open_time"`
		RefundPolicy struct {
			Days      int    `json:"days"`
			Notes     string `json:"notes"`
			Policiies string `json:"policies"`
		} `json:"refund_policy"`
		Waitlisting string `json:"wait_listing"`
	} `json:"rsvp_rules"`
	RSVPSample struct {
		Created MeetupTime `json:"created"`
		ID      int        `json:"id"`
		Member  struct {
			Bio          string `json:"bio"`
			EventContext string `json:"event_context"`
			ID           int    `json:"id"`
			Name         string `json:"name"`
			Photo        string `json:"photo"`
			Role         string `json:"role"`
			Self         struct {
				Actions   string `json:"self"`
				PayStatus string `json:"pay_status"`
				Role      string `json:"role"`
				RSVP      struct {
					Answers  []string `json:"answers"`
					Guests   int      `json:"guests"`
					Response string   `json:"response"`
				} `json:"rsvp"`
			} `json:"self"`
			Title string `json:"title"`
		} `json:"member"`
		Updated MeetupTime `json:"updated"`
	} `json:"rsvp_sample"`
	RSVPAble          bool `json:"rsvpable"`
	RSVPAbleAfterJoin bool `json:"rsvpable_after_join"`
	Saved             bool `json:"saved"`
	Self              struct {
		Actions   string `json:"self"`
		PayStatus string `json:"pay_status"`
		Role      string `json:"role"`
		RSVP      struct {
			Answers  []string `json:"answers"`
			Guests   int      `json:"guests"`
			Response string   `json:"response"`
		} `json:"rsvp"`
	} `json:"self"`
	Series struct {
		Description string     `json:"description"`
		EndDate     MeetupTime `json:"end_date"`
		ID          int        `json:"id"`
		Monthly     struct {
			DayOfWeek   int `json:"day_of_week"`
			Interval    int `json:"interval"`
			WeekOfMonth int `json:"week_of_month"`
		} `json:"monthly"`
		StartDate       MeetupTime `json:"start_date"`
		TemplateEventID int        `json:"template_event_id"`
		Weekly          struct {
			DayOfWeek []int `json:"days_of_week"`
			Interval  int   `json:"interval"`
		} `json:"weekly"`
	} `json:"series"`
	ShortLink            string `json:"short_link"`
	SimpeHTMLDescription string `json:"simple_html_description"`
	Status               string `json:"status"`
	SurveyQuestions      []struct {
		ID       int    `json:"id"`
		Question string `json:"question"`
	} `json:"survey_questions"`
	Time      MeetupTime `json:"time"`
	Updated   MeetupTime `json:"updated_time"`
	UTCOffset int64      `json:"utc_offset"`
	Venue     struct {
		Address1             string  `json:"address_1"`
		Address2             string  `json:"address_2"`
		Address3             string  `json:"address_3"`
		City                 string  `json:"city"`
		Country              string  `json:"country"`
		ID                   int     `json:"id"`
		Lat                  float64 `json:"lat"`
		Lon                  float64 `json:"lon"`
		LocalizedCountryName string  `json:"localized_country_name"`
		Name                 string  `json:"Name"`
		Phone                string  `json:"phone"`
		Repinned             bool    `json:"repinned"`
		State                string  `json:"state"`
		Zip                  string  `json:"zip"`
	} `json:"venue"`
	VenueVisibility string `json:"venue_visibility"`
	Visibility      string `json:"visibility"`
	WaitlistCount   int    `json:"waitlist_count"`
	Why             string `json:"why"`
	YesRSVPCount    int    `json:"yes_rsvp_count"`
}
