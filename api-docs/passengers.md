# /api/passenger

## Actions
**1. Create a Passenger** 
|Method| Route |
|--|--|
| POST | /api/passenger |

    {   
	    "first_name": "Jamiu",
	    "middle_name": "Jonathan",
	    "last_name": "Adeko",
	    "passport_number": "A60F6B2",
	    "age": 45
    }   

**2. List All Registered Passengers**
|Method|Route  |
|--|--|
| GET |/api/passenger  |

**3. Get Passenger by ID**
|Method|Route  |
|--|--|
|  GET|/api/passenger/{ID}  |

**4. Get Passenger by User id**
|Method|Route  |
|--|--|
| GET |  /api/passenger/user/{user id}|

**5. Delete Passenger by ID**
|Method|Route  |
|--|--|
|DELETE  | /api/passenger/{ID} |



