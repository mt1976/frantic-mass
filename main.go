package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/baseline"
	"github.com/mt1976/frantic-mass/app/dao/dateIndex"
	"github.com/mt1976/frantic-mass/app/dao/goal"
	"github.com/mt1976/frantic-mass/app/dao/tag"
	user "github.com/mt1976/frantic-mass/app/dao/user"
	"github.com/mt1976/frantic-mass/app/dao/weight"
	"github.com/mt1976/frantic-mass/app/dao/weightTag"
	"github.com/mt1976/frantic-mass/app/functions"
	"github.com/mt1976/frantic-mass/app/jobs"
	"github.com/mt1976/frantic-mass/app/types/measures"
	"github.com/mt1976/frantic-mass/app/web/contentProvider"
	endpointprovider "github.com/mt1976/frantic-mass/app/web/endpointProvider"
	"github.com/mt1976/frantic-mass/app/web/handlers"
	my_middleware "github.com/mt1976/frantic-mass/app/web/middleware"
)

func main() {
	// This is the main function

	logHandler.InfoLogger.Println("Start")
	// Initialise the database connections

	common, _ := functions.GetConfig()

	logHandler.InfoLogger.Println("Initialising Database Connections")

	err := functions.Initialise(context.TODO())
	if err != nil {
		logHandler.ErrorLogger.Println("Error during initialisation:", err)
		return
	}

	logHandler.InfoLogger.Println("Database Connections Initialised Successfully")

	logHandler.InfoLogger.Println("Reset Database, clearing down existing records")
	err = functions.ClearDown(context.TODO())
	if err != nil {
		logHandler.ErrorLogger.Println("Error during clear down:", err)
		return
	}

	logHandler.InfoLogger.Println("Database Cleared Successfully")

	logHandler.InfoLogger.Println("Creating User and related records")
	// Create a user and related records
	// This is a test user, you can change the ID to create multiple users
	// The user ID will be used for the rest of the operations
	// You can also change the weight and goal values as per your requirements

	//currentWeight := types.NewWeight(114.0)

	userIdentifier := 100 // This is the user ID we will use for the rest of the operations
	logHandler.InfoLogger.Println("Creating User with ID:", userIdentifier)
	//	for i := 0; i < 1; i++ {
	//		randNum := rand.Intn(10000-1000) + 1000
	//logHandler.InfoLogger.Println("randNum:", randNum)

	thisUser, newRedErr := user.Create(context.TODO(), fmt.Sprintf("user %v", userIdentifier), fmt.Sprintf("user_%v", userIdentifier), fmt.Sprintf("password_%v", userIdentifier), fmt.Sprintf("user_%v@example.com", userIdentifier))
	if newRedErr != nil {
		logHandler.ErrorLogger.Println(newRedErr)
	}

	userIdentifier = thisUser.ID
	logHandler.InfoLogger.Printf("User Created:[%+v]", thisUser)

	thisUser2, newRedErr := user.Create(context.TODO(), fmt.Sprintf("user %v", userIdentifier+1), fmt.Sprintf("user_%v", userIdentifier+1), fmt.Sprintf("password_%v", userIdentifier+1), fmt.Sprintf("user_%v@example.com", userIdentifier+1))
	if newRedErr != nil {
		logHandler.ErrorLogger.Println(newRedErr)
	}

	logHandler.InfoLogger.Printf("User Created:[%+v]", thisUser2)

	//	}

	// user.ExportRecordsAsCSV()
	// // Drop Again
	// dropErr = user.ClearDown(context.TODO())
	// if dropErr != nil {
	// 	logHandler.ErrorLogger.Println(dropErr)
	// } else {
	// 	logHandler.InfoLogger.Printf("Drop Data Post Export:[templ]")
	// }

	// user.ImportRecordsFromCSV()

	cdrop, cdroperr := user.Count()
	if cdroperr != nil {
		logHandler.ErrorLogger.Println(cdroperr)
	} else {
		logHandler.InfoLogger.Printf("Post Import:[%+v]", cdrop)
	}

	//Set Date of Birth to 27/02/1976
	dob := time.Date(1976, 2, 27, 0, 0, 0, 0, time.UTC)
	logHandler.InfoLogger.Println("Setting Date of Birth for UserID:", userIdentifier, "to", dob)

	logHandler.InfoLogger.Println("Creating Baseline for UserID:", userIdentifier)
	thisBaseline, baselineErr := baseline.Create(context.TODO(), userIdentifier, measures.Height{CMs: 187.96}, 6, fmt.Sprintf("BaselineFor%v", userIdentifier), dob, time.Now())
	if baselineErr != nil {
		logHandler.ErrorLogger.Println(baselineErr)
	} else {
		logHandler.InfoLogger.Printf("Baseline Created:[%+v]", thisBaseline)
	}

	logHandler.InfoLogger.Println("Creating Baseline for UserID:", thisUser2.ID)
	thisBaseline2, baselineErr2 := baseline.Create(context.TODO(), thisUser2.ID, measures.Height{CMs: 170.96}, 6, fmt.Sprintf("BaselineFor%v", thisUser2.ID), dob, time.Now())
	if baselineErr2 != nil {
		logHandler.ErrorLogger.Println(baselineErr2)
	} else {
		logHandler.InfoLogger.Printf("Baseline Created:[%+v]", thisBaseline2)
	}

	logHandler.InfoLogger.Println("Creating Goal for UserID:", userIdentifier)
	thisGoal, goalErr := goal.Create(context.TODO(), userIdentifier, fmt.Sprintf("GoalFor%v", userIdentifier), measures.Weight{KGs: 90.00}, time.Now().AddDate(0, 0, 30), measures.Weight{KGs: 2.0}, "This is a test goal to check the goal creation process", false)

	if goalErr != nil {
		logHandler.ErrorLogger.Println(goalErr)
	} else {
		logHandler.InfoLogger.Printf("Goal Created:[%+v]", thisGoal)
	}

	logHandler.InfoLogger.Println("Creating Goal for UserID:", userIdentifier)
	thisGoal2, goalErr2 := goal.Create(context.TODO(), userIdentifier, fmt.Sprintf("GoalFor%v2", userIdentifier), measures.Weight{KGs: 86.00}, time.Now().AddDate(0, 0, 30), measures.Weight{KGs: 2.0}, "This is a test goal to check the goal creation process", false)

	if goalErr2 != nil {
		logHandler.ErrorLogger.Println(goalErr2)
	} else {
		logHandler.InfoLogger.Printf("Goal Created:[%+v]", thisGoal2)
	}

	logHandler.InfoLogger.Println("Creating Goal for UserID:", userIdentifier)
	thisGoal3, goalErr3 := goal.Create(context.TODO(), userIdentifier, fmt.Sprintf("GoalFor%v3", userIdentifier), measures.Weight{KGs: 84.78950784000001}, time.Now().AddDate(0, 0, 30), measures.Weight{KGs: 2.0}, "This is a test goal to check the goal creation process", false)

	if goalErr3 != nil {
		logHandler.ErrorLogger.Println(goalErr3)
	} else {
		logHandler.InfoLogger.Printf("Goal Created:[%+v]", thisGoal3)
	}

	//84.78950784000001

	avgGoal, avgGoalErr := goal.Create(context.TODO(), userIdentifier, fmt.Sprintf("AvgGoalFor%v", userIdentifier), measures.Weight{KGs: 90.00}, time.Now().AddDate(0, 0, 30), measures.Weight{KGs: 0}, "This is an average weight loss goal", true)
	if avgGoalErr != nil {
		logHandler.ErrorLogger.Println(avgGoalErr)
	} else {
		logHandler.InfoLogger.Printf("Average Goal Created:[%+v]", avgGoal)
	}

	logHandler.InfoLogger.Println("Creating Tag")
	thisTag, tagErr := tag.Create(context.TODO(), fmt.Sprintf("TagFor%v", userIdentifier))
	if tagErr != nil {
		logHandler.ErrorLogger.Println(tagErr)
	} else {
		logHandler.InfoLogger.Printf("Tag Created:[%+v]", thisTag)
	}

	// Import a CSV file of Weight records
	logHandler.InfoLogger.Println("Importing Weight Records from CSV")
	err = weight.ImportRecordsFromCSV()
	if err != nil {
		logHandler.ErrorLogger.Println("Error importing weight records from CSV:", err)
	} else {
		logHandler.InfoLogger.Println("Weight records imported successfully from CSV")
	}

	// logHandler.InfoLogger.Println("Creating Weight for UserID:", userIdentifier)
	// thisWeight, weightErr := weight.Create(context.TODO(), userIdentifier, *currentWeight, fmt.Sprintf("WeightFor %v", userIdentifier), time.Now())
	// if weightErr != nil {
	// 	logHandler.ErrorLogger.Println(weightErr)
	// } else {
	// 	logHandler.InfoLogger.Printf("Weight Created:[%+v]", thisWeight)
	// }

	// CREATE A today DATEINDEX

	var dij jobs.DateIndexJob = jobs.DateIndexJob{}
	dij.AddDatabaseAccessFunctions(weight.GetDatabaseConnections())
	logHandler.InfoLogger.Println("Running DateIndex Job")
	err = dij.Run()

	if err != nil {
		logHandler.ErrorLogger.Println("Error creating DateIndex:", err)
	} else {
		logHandler.InfoLogger.Printf("DateIndexs created successfully")
	}

	//_ = dateIndex.ExportRecordsAsCSV()

	tdID, td, tdErr := dateIndex.GetToday()
	if tdErr != nil {
		logHandler.ErrorLogger.Println("Error getting today's DateIndex:", tdErr)
	} else {
		logHandler.InfoLogger.Printf("Today's DateIndex: %+v [%v]", td, tdID)
	}

	//os.Exit(0) // Exit the program successfully
	thisWeight, err := functions.FetchLatestWeightRecord(thisUser.ID)
	if err != nil {
		logHandler.ErrorLogger.Println("Error fetching latest weight record:", err)
	} else {
		logHandler.InfoLogger.Printf("Latest Weight Record for UserID %d: %+v", thisUser.ID, thisWeight)
	}

	// Rebuild the weight projection
	logHandler.InfoLogger.Println("Rebuilding Weight Projections for UserID:", userIdentifier)
	err = functions.BuildWeightGoalsProjections(thisUser, thisWeight.Weight)
	if err != nil {
		logHandler.ErrorLogger.Println("Error rebuilding weight projections:", err)
	} else {
		logHandler.InfoLogger.Println("Weight projection rebuilt successfully")
	}
	logHandler.InfoLogger.Println("Creating WeightTag for UserID:", userIdentifier)

	logHandler.InfoLogger.Println("Creating WeightTag")
	thisWeightTag, weightTagErr := weightTag.Create(context.TODO(), thisWeight.ID, thisTag.ID)
	if weightTagErr != nil {
		logHandler.ErrorLogger.Println(weightTagErr)
	} else {
		logHandler.InfoLogger.Printf("WeightTag Created:[%+v]", thisWeightTag)
	}

	logHandler.InfoLogger.Println("HeightCM", thisBaseline.Height.CmAsString())
	logHandler.InfoLogger.Println("HeightM", thisBaseline.Height.MetersAsString())

	logHandler.InfoLogger.Println("HeightIN", thisBaseline.Height.InchesAsString())
	logHandler.InfoLogger.Println("HeightFT", thisBaseline.Height.FeetAsString())

	logHandler.InfoLogger.Println("HeightString", thisBaseline.Height.String())

	logHandler.InfoLogger.Println("WeightKg", thisWeight.Weight.KgAsString())
	logHandler.InfoLogger.Println("WeightLbs", thisWeight.Weight.LbsAsString())
	logHandler.InfoLogger.Println("WeightStone", thisWeight.Weight.StonesAsString())
	logHandler.InfoLogger.Println("WeightString", thisWeight.Weight.String())

	avg, tot, lossErr := functions.AverageWeightLoss(userIdentifier)
	if lossErr != nil {
		logHandler.ErrorLogger.Println(lossErr)
	} else {
		logHandler.InfoLogger.Printf("Average Weight Loss for User %d: %s", userIdentifier, avg.KgAsString())
		logHandler.InfoLogger.Printf("Total Weight Loss for User %d: %s", userIdentifier, tot.KgAsString())
	}

	// logHandler.InfoLogger.Println("Exporting Data")
	// err = functions.ExportDataSnapshot()
	// if err != nil {
	// 	logHandler.ErrorLogger.Println("Error exporting data:", err)
	// } else {
	// 	logHandler.InfoLogger.Println("Data exported successfully")
	// }

	logHandler.InfoLogger.Println("Starting Jobs")
	err = functions.StartJobs(context.TODO(), common)
	if err != nil {
		logHandler.ErrorLogger.Println("Error starting jobs:", err)
	} else {
		logHandler.InfoLogger.Println("Jobs started successfully")
	}

	logHandler.InfoLogger.Println("All operations completed successfully")

	// uv, err := viewProvider.Users(context.TODO())
	// if err != nil {
	// 	logHandler.ErrorLogger.Println("Error creating UserChooser view:", err)
	// } else {
	// 	logHandler.EventLogger.Println("UserChooser view created successfully")
	// }

	// godump.Dump(uv)

	logHandler.InfoLogger.Println("End of main function")

	r := chi.NewRouter()

	// Create the HTTP server
	server := &http.Server{
		Addr:    ":" + common.GetServer_PortString(),
		Handler: r,
	}

	r.Use(middleware.Logger)
	logHandler.InfoLogger.Println("Using Logger Middleware")
	r.Use(middleware.RequestID)
	logHandler.InfoLogger.Println("Using RequestID Middleware")
	r.Use(middleware.RealIP)
	logHandler.InfoLogger.Println("Using RealIP Middleware")
	r.Use(middleware.Recoverer)
	logHandler.InfoLogger.Println("Using Recoverer Middleware")
	r.Use(my_middleware.HandleHTTPMethodConversion)
	logHandler.InfoLogger.Println("Using HandleHTTPMethodConversion Middleware")
	if common.GetServer_Compression() {
		//r.Use(middleware.Compress(5, "gzip"))
		r.Use(my_middleware.HandleHTMLMinification())
		logHandler.InfoLogger.Println("Using HTML Minification Middleware")
		// // /* means to compress all content types that can be compressed.
		// compressor := middleware.NewCompressor(5, "/*")
		// compressor.SetEncoder("br", func(w io.Writer, level int) io.Writer {
		// 	params := brotli_enc.NewBrotliParams()
		// 	params.SetQuality(level)
		// 	return brotli_enc.NewBrotliWriter(params, w)
		// })
		// r.Use(compressor.Handler)
		r.Use(my_middleware.HandleBrotli)
		logHandler.InfoLogger.Println("Using Brotli Compression Middleware")
		logHandler.InfoLogger.Println("Compression is enabled, using Brotli and HTML Minification, all responses will be compressed")
	} else {
		logHandler.InfoLogger.Println("Compression is disabled")
	}

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))
	logHandler.InfoLogger.Println("Using Timeout Middleware with 60 seconds")
	// r.Use(middleware.StripSlashes)
	// logHandler.InfoLogger.Println("Using StripSlashes Middleware")
	r.Use(middleware.URLFormat)
	logHandler.InfoLogger.Println("Using URLFormat Middleware")
	r.Get(contentProvider.LauncherURI, handlers.Launcher)
	r.Get(contentProvider.UserChooserURI, handlers.UserChooser)
	r.Get(contentProvider.DashboardURI, handlers.Dashboard)                       // Placeholder for user dashboard handler
	r.Get(contentProvider.UserURI, handlers.UserRead)                             // View/Edit handler
	r.Post(contentProvider.UserURI, handlers.UserCreate)                          // New User handler
	r.Put(contentProvider.UserURI, handlers.UserUpdate)                           // Update User handler
	r.Get(contentProvider.GoalURI, handlers.Goal)                                 // Placeholder for goal edit handlerx
	r.Get(contentProvider.ProjectionURI, handlers.Projection)                     // Placeholder for projection handler
	r.Get(contentProvider.WeightURI, handlers.Weight)                             // Placeholder for weight edit handler
	r.Get(contentProvider.TestURI, handlers.Test)                                 // Placeholder for test handler
	r.Get(endpointprovider.BMIUserWeightEndpoint, endpointprovider.BMI)           // BMI calculation endpoint\
	r.Get(endpointprovider.BMIEnrichmentEndpoint, endpointprovider.BMIEnrichment) // BMI enrichment endpoint
	r.Get(endpointprovider.BMIWeightEndpoint, endpointprovider.WeightFromBMI)     // Weight from BMI endpoint
	//r.Get(contentProvider.WeightURI, handlers.Weight)         // Placeholder for weight edit handler
	r.Get("/test", handlers.Dummy)
	// Inject shutdown function into handler
	r.Handle("/shutdown", handlers.ShutdownHandler(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			logHandler.ErrorLogger.Println("Error during shutdown:", err)
		}
	}))
	r.NotFound(handlers.NotFound)
	r.MethodNotAllowed(handlers.MethodNotAllowed)
	//r.Handle("/favicon.ico", http.FileServer(http.Dir("./res/images")))
	r.Handle("/my.css/*", http.StripPrefix("/my.css/", http.FileServer(http.Dir("./res/css"))))
	r.Handle("/my.js/*", http.StripPrefix("/my.js/", http.FileServer(http.Dir("./res/js"))))
	r.Handle("/my.img/*", http.StripPrefix("/my.img/", http.FileServer(http.Dir("./res/images"))))
	r.Handle("/pico.css/*", http.StripPrefix("/pico.css/", http.FileServer(http.Dir("./node_modules/@picocss/pico/css"))))
	r.Handle("/pico.js/*", http.StripPrefix("/pico.js/", http.FileServer(http.Dir("./node_modules/@picocss/pico/js"))))
	//r.Handle("/favicon.ico", http.FileServer(http.Dir("./res/images")))
	r.Get("/manifest", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./res/app/manifest.json")
	})

	r.Handle("/glyphs/*", http.StripPrefix("/glyphs/", http.FileServer(http.Dir("./node_modules/bootstrap-icons/font"))))
	r.Handle("/images/*", http.StripPrefix("/images/", http.FileServer(http.Dir("./res/images"))))
	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./res/images/favicon.ico")
	})
	r.Get("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./res/robots.txt")
	})
	r.NotFound(handlers.NotFound)
	r.MethodNotAllowed(handlers.MethodNotAllowed)
	//r.Handle("/favicon.ico", http.FileServer(http.Dir("./res/images")))

	logHandler.InfoLogger.Println("Starting server on port", common.GetServer_Port())
	// Start the server
	http.Handle("/", r)
	// Use the port from the configuration

	// Shutdown on SIGINT
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)
		<-stop
		logHandler.WarningLogger.Println("Interrupt received. Shutting down...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(ctx)
		os.Exit(0) // Exit the program successfully
		logHandler.InfoLogger.Println("Server shut down cleanly")
	}()

	// Walk through all registered routes and print them
	err = chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("Registered route: %8s %30s %v\n", method, route, getHandlerName(handler))
		return nil
	})
	if err != nil {
		fmt.Printf("Failed to walk routes: %v\n", err)
	}

	logHandler.InfoLogger.Printf("Server is running at http://127.0.0.1:%v", common.GetServer_PortString())

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logHandler.ErrorLogger.Println("Server error:", err)
	}
	logHandler.InfoLogger.Println("Server shut down cleanly")

	// If you want to use a specific port, you can uncomment the line below
	//http.ListenAndServe(":3000", r)
	//os.Exit(0) // Exit the program successfully
}

func getHandlerName(h http.Handler) string {
	if handlerFunc, ok := h.(http.HandlerFunc); ok {
		return runtime.FuncForPC(reflect.ValueOf(handlerFunc).Pointer()).Name()
	}
	return fmt.Sprintf("%T", h)
}
