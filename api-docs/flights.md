

# /api/flight

## Actions

**1. Create a Flight Entry**

|Method|Route  |
|--|--|
|POST  |/api/flight  |


    {   
	    "name": "Air Peace",
	    "tail_number": "47-HJU",
	    "flight_number": 12430,
	    "departure_city": "ABJ",
	    "destination_city": "KAN",
	    "seats": 110,
	    "seats_left": 37,
	    "business_class_price": 22200,
	    "economy_class_price": 145500  
	    "flight_date":  "2024-11-26T01:00:00+01:00"
    }  

**2.  Find Flights by Departure and Destination city**
|Method|Route  |
|--|--|
|  POST|/api/flight/find  |

    {   
	    "departure_city": "LAG",
	    "destination_city": "KAN"
    }


**3. Get Flights  by Flight Number**
    
|Method|Route  |
|--|--|
|  GET|/api/flight/{flight number}  |
**4. Get Flight by Flight Id**
|Method|Route  |
|--|--|
|  GET|/api/flight/id/{flight id}  |


**5. Get All Flights entries**
|Method|Route  |
|--|--|
| GET | /api/flight |

**6. Get Flights by Date**
    
|Method|Route  |
|--|--|
|POST  |/api/flight/date  |

       
    {"flight_date": "2024-11-26" }

 
