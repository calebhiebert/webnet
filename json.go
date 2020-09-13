package web_net

import "encoding/json"

const JSONData = `[
  {
    "_id": "5f5830a10687b03872e8bc1b",
    "index": 0,
    "guid": "8330155c-8cc1-4657-b63c-257923c9f709",
    "isActive": true,
    "balance": "$3,624.63",
    "picture": "http://placehold.it/32x32",
    "age": 30,
    "eyeColor": "green",
    "name": {
      "first": "White",
      "last": "Cook"
    },
    "company": "GAZAK",
    "email": "white.cook@gazak.biz",
    "phone": "+1 (983) 572-2582",
    "address": "463 Debevoise Avenue, Trucksville, Utah, 2769",
    "about": "Incididunt do consectetur cillum qui magna laborum occaecat aliquip. Officia cupidatat incididunt pariatur cillum qui commodo culpa reprehenderit exercitation elit deserunt. Id sunt sunt dolor incididunt in magna eu magna do elit aliquip velit reprehenderit. Consequat duis ipsum sit ad ad ut aliqua veniam tempor nulla eiusmod laborum. Commodo incididunt labore elit laborum ipsum occaecat cillum amet adipisicing nisi aute.",
    "registered": "Friday, October 17, 2014 5:18 PM",
    "latitude": "10.017589",
    "longitude": "116.622991",
    "tags": [
      "tempor",
      "dolor",
      "enim",
      "nisi",
      "deserunt"
    ],
    "range": [
      0,
      1,
      2,
      3,
      4,
      5,
      6,
      7,
      8,
      9
    ],
    "friends": [
      {
        "id": 0,
        "name": "Young Hull"
      },
      {
        "id": 1,
        "name": "Maura Goodman"
      },
      {
        "id": 2,
        "name": "Burns Stokes"
      }
    ],
    "greeting": "Hello, White! You have 10 unread messages.",
    "favoriteFruit": "strawberry"
  },
  {
    "_id": "5f5830a1c747fc5cb43a5252",
    "index": 1,
    "guid": "486bfa76-cfe8-4df9-abd3-5a01b75e6286",
    "isActive": false,
    "balance": "$2,472.66",
    "picture": "http://placehold.it/32x32",
    "age": 33,
    "eyeColor": "brown",
    "name": {
      "first": "Roxie",
      "last": "Hampton"
    },
    "company": "EXOSIS",
    "email": "roxie.hampton@exosis.com",
    "phone": "+1 (904) 562-2707",
    "address": "833 Lawrence Street, Kieler, Nevada, 1049",
    "about": "Labore ad ex voluptate nostrud ex laboris non laboris sunt dolore occaecat officia nisi minim. Adipisicing irure est minim minim. Pariatur non pariatur amet nostrud nostrud. Elit excepteur aliquip sunt dolor. Sunt amet veniam et voluptate laboris excepteur nisi ex tempor esse quis qui deserunt. Culpa elit anim non fugiat eu nulla laboris aliqua exercitation. Laborum officia dolore exercitation adipisicing tempor irure sunt cillum incididunt in.",
    "registered": "Saturday, March 22, 2014 3:45 PM",
    "latitude": "-80.43566",
    "longitude": "117.27544",
    "tags": [
      "laboris",
      "velit",
      "irure",
      "laboris",
      "excepteur"
    ],
    "range": [
      0,
      1,
      2,
      3,
      4,
      5,
      6,
      7,
      8,
      9
    ],
    "friends": [
      {
        "id": 0,
        "name": "Brianna Watson"
      },
      {
        "id": 1,
        "name": "Noreen Woodward"
      },
      {
        "id": 2,
        "name": "Stacie Johns"
      }
    ],
    "greeting": "Hello, Roxie! You have 9 unread messages.",
    "favoriteFruit": "banana"
  },
  {
    "_id": "5f5830a1d65ca97962c9d5b5",
    "index": 2,
    "guid": "df3ece00-2a55-4b78-bebc-1c56b012a8e0",
    "isActive": false,
    "balance": "$3,280.31",
    "picture": "http://placehold.it/32x32",
    "age": 20,
    "eyeColor": "blue",
    "name": {
      "first": "Wall",
      "last": "Lancaster"
    },
    "company": "HANDSHAKE",
    "email": "wall.lancaster@handshake.ca",
    "phone": "+1 (972) 420-3407",
    "address": "108 Ferry Place, Epworth, West Virginia, 5970",
    "about": "Mollit aliqua officia cillum Lorem eu ipsum ipsum labore anim proident qui pariatur enim. Mollit dolor laboris esse consectetur pariatur consectetur proident non nostrud eiusmod elit. In occaecat excepteur nisi ad excepteur nostrud incididunt sunt anim occaecat sit magna quis. Sint sunt excepteur mollit excepteur occaecat ullamco eiusmod dolor cillum tempor et Lorem. Sint commodo laboris nisi laborum adipisicing dolore in exercitation. Do nostrud minim non exercitation. Duis quis ullamco voluptate pariatur sit elit deserunt dolore.",
    "registered": "Sunday, June 16, 2019 9:59 AM",
    "latitude": "87.860956",
    "longitude": "74.936819",
    "tags": [
      "eiusmod",
      "do",
      "mollit",
      "laborum",
      "pariatur"
    ],
    "range": [
      0,
      1,
      2,
      3,
      4,
      5,
      6,
      7,
      8,
      9
    ],
    "friends": [
      {
        "id": 0,
        "name": "Muriel Hood"
      },
      {
        "id": 1,
        "name": "Ashley Trujillo"
      },
      {
        "id": 2,
        "name": "Davis Collins"
      }
    ],
    "greeting": "Hello, Wall! You have 9 unread messages.",
    "favoriteFruit": "strawberry"
  },
  {
    "_id": "5f5830a17646558f8743dc73",
    "index": 3,
    "guid": "c7c30f87-600e-451e-9ac7-daa02251080e",
    "isActive": false,
    "balance": "$1,530.45",
    "picture": "http://placehold.it/32x32",
    "age": 32,
    "eyeColor": "brown",
    "name": {
      "first": "Flowers",
      "last": "Abbott"
    },
    "company": "COMTRACT",
    "email": "flowers.abbott@comtract.co.uk",
    "phone": "+1 (872) 486-2612",
    "address": "673 Bergen Street, Wintersburg, New Mexico, 1994",
    "about": "Eiusmod cillum eu aute eiusmod ipsum ullamco. Duis voluptate consectetur culpa in in et dolore irure sit laboris aliqua et. Nulla quis ea anim eiusmod non et pariatur laboris officia. Aute magna culpa magna eu sint et nisi.",
    "registered": "Thursday, January 23, 2014 6:12 PM",
    "latitude": "78.700956",
    "longitude": "120.752362",
    "tags": [
      "eiusmod",
      "non",
      "commodo",
      "reprehenderit",
      "id"
    ],
    "range": [
      0,
      1,
      2,
      3,
      4,
      5,
      6,
      7,
      8,
      9
    ],
    "friends": [
      {
        "id": 0,
        "name": "Morse Le"
      },
      {
        "id": 1,
        "name": "Magdalena Reyes"
      },
      {
        "id": 2,
        "name": "Maldonado Gray"
      }
    ],
    "greeting": "Hello, Flowers! You have 7 unread messages.",
    "favoriteFruit": "apple"
  },
  {
    "_id": "5f5830a1e12f937da6a1c728",
    "index": 4,
    "guid": "9f1a9248-97f6-4e75-8fda-668da6694059",
    "isActive": true,
    "balance": "$3,492.33",
    "picture": "http://placehold.it/32x32",
    "age": 37,
    "eyeColor": "brown",
    "name": {
      "first": "Blevins",
      "last": "Woods"
    },
    "company": "RUGSTARS",
    "email": "blevins.woods@rugstars.io",
    "phone": "+1 (870) 488-3919",
    "address": "762 Montgomery Place, Sardis, Tennessee, 4800",
    "about": "Esse ut aute duis mollit. Esse tempor eiusmod nostrud sint amet fugiat fugiat culpa ad incididunt ad. Cillum incididunt incididunt laboris non pariatur amet et adipisicing adipisicing fugiat ut cillum. Reprehenderit dolor ullamco occaecat ipsum qui et. Minim sint non amet officia culpa. Eu eiusmod Lorem eu aute cillum commodo aliquip aliquip qui. Elit Lorem adipisicing exercitation nisi anim sit deserunt culpa proident consequat cupidatat cupidatat amet pariatur.",
    "registered": "Monday, September 30, 2019 7:11 PM",
    "latitude": "-60.846886",
    "longitude": "-78.994875",
    "tags": [
      "duis",
      "veniam",
      "incididunt",
      "proident",
      "proident"
    ],
    "range": [
      0,
      1,
      2,
      3,
      4,
      5,
      6,
      7,
      8,
      9
    ],
    "friends": [
      {
        "id": 0,
        "name": "Morgan Stanton"
      },
      {
        "id": 1,
        "name": "Pearson Fischer"
      },
      {
        "id": 2,
        "name": "Roach Skinner"
      }
    ],
    "greeting": "Hello, Blevins! You have 5 unread messages.",
    "favoriteFruit": "apple"
  },
  {
    "_id": "5f5830a117fbe9268e20244f",
    "index": 5,
    "guid": "00f6348c-3cc8-4b6e-b7ad-6ac425d0cb19",
    "isActive": false,
    "balance": "$3,809.17",
    "picture": "http://placehold.it/32x32",
    "age": 24,
    "eyeColor": "blue",
    "name": {
      "first": "Angelica",
      "last": "Maldonado"
    },
    "company": "MEDESIGN",
    "email": "angelica.maldonado@medesign.info",
    "phone": "+1 (936) 565-3922",
    "address": "499 Kensington Walk, Grenelefe, Delaware, 2200",
    "about": "Occaecat ipsum labore velit mollit sit labore ipsum. Commodo Lorem dolor incididunt eu sit aliquip minim exercitation eiusmod laboris commodo quis do. Consectetur commodo veniam consectetur ipsum duis aliqua non consequat sint officia enim id. Exercitation voluptate adipisicing fugiat labore officia occaecat ullamco fugiat ex. Lorem occaecat est do sunt pariatur veniam.",
    "registered": "Saturday, May 9, 2015 2:27 PM",
    "latitude": "34.07016",
    "longitude": "-96.727205",
    "tags": [
      "sit",
      "adipisicing",
      "aliquip",
      "id",
      "culpa"
    ],
    "range": [
      0,
      1,
      2,
      3,
      4,
      5,
      6,
      7,
      8,
      9
    ],
    "friends": [
      {
        "id": 0,
        "name": "Keller House"
      },
      {
        "id": 1,
        "name": "King Soto"
      },
      {
        "id": 2,
        "name": "Salas Alford"
      }
    ],
    "greeting": "Hello, Angelica! You have 7 unread messages.",
    "favoriteFruit": "banana"
  }
]`

const JSONData2 = `[
  {
    "_id": "5f5830a10687b03872e8bc1b",
    "index": 0,
    "guid": "8330155c-8cc1-4657-b63c-257923c9f709",
    "isActive": true,
    "balance": "$3,624.63",
    "picture": "http://placehold.it/32x32",
    "age": 30,
    "eyeColor": "green",
    "name": {
      "first": "White",
      "last": "Cook"
    },
    "company": "GAZAK",
    "email": "white.cook@gazak.biz",
    "phone": "+1 (983) 572-2582",
    "address": "463 Debevoise Avenue, Trucksville, Utah, 2769",
    "about": "Incididunt do consectetur cillum qui magna laborum occaecat aliquip. Officia cupidatat incididunt pariatur cillum qui commodo culpa reprehenderit exercitation elit deserunt. Id sunt sunt dolor incididunt in magna eu magna do elit aliquip velit reprehenderit. Consequat duis ipsum sit ad ad ut aliqua veniam tempor nulla eiusmod laborum. Commodo incididunt labore elit laborum ipsum occaecat cillum amet adipisicing nisi aute.",
    "registered": "Friday, October 17, 2014 5:18 PM",
    "latitude": "10.017589",
    "longitude": "116.622991",
    "tags": [
      "tempor",
      "dolor",
      "enim",
      "nisi",
      "deserunt"
    ],
    "range": [
      0,
      1,
      2,
      3,
      4,
      5,
      6,
      7,
      8,
      9
    ],
    "friends": [
      {
        "id": 0,
        "name": "Young Hull"
      },
      {
        "id": 1,
        "name": "Maura Goodman"
      },
      {
        "id": 2,
        "name": "Burns Stokes"
      }
    ],
    "greeting": "Hello, White! You have 10 unread messages.",
    "favoriteFruit": "strawberry"
  },
  {
    "_id": "5f5830a1c747fc5cb43a5252",
    "index": 1,
    "guid": "486bfa76-cfe8-4df9-abd3-5a01b75e6286",
    "isActive": false,
    "balance": "$2,472.66",
    "picture": "http://placehold.it/32x32",
    "age": 33,
    "eyeColor": "brown",
    "name": {
      "first": "Roxie",
      "last": "Hampton"
    },
    "company": "EXOSIS",
    "email": "roxie.hampton@exosis.com",
    "phone": "+1 (904) 562-2707",
    "address": "833 Lawrence Street, Kieler, Nevada, 1049",
    "about": "Labore ad ex voluptate nostrud ex laboris non laboris sunt dolore occaecat officia nisi minim. Adipisicing irure est minim minim. Pariatur non pariatur amet nostrud nostrud. Elit excepteur aliquip sunt dolor. Sunt amet veniam et voluptate laboris excepteur nisi ex tempor esse quis qui deserunt. Culpa elit anim non fugiat eu nulla laboris aliqua exercitation. Laborum officia dolore exercitation adipisicing tempor irure sunt cillum incididunt in.",
    "registered": "Saturday, March 22, 2014 3:45 PM",
    "latitude": "-80.43566",
    "longitude": "117.27544",
    "tags": [
      "laboris",
      "velit",
      "irure",
      "laboris",
      "excepteur"
    ],
    "range": [
      0,
      1,
      2,
      3,
      4,
      5,
      6,
      7,
      8,
      9
    ],
    "friends": [
      {
        "id": 0,
        "name": "Brianna Watson"
      },
      {
        "id": 1,
        "name": "Noreen Woodward"
      },
      {
        "id": 2,
        "name": "Stacie Johns"
      }
    ],
    "greeting": "Hello, Roxie! You have 9 unread messages.",
    "favoriteFruit": "banana"
  },
  {
    "_id": "5f5830a1d65ca97962c9d5b5",
    "index": 2,
    "guid": "df3ece00-2a55-4b78-bebc-1c56b012a8e0",
    "isActive": false,
    "balance": "$3,280.31",
    "picture": "http://placehold.it/32x32",
    "age": 20,
    "eyeColor": "blue",
    "name": {
      "first": "Wall",
      "last": "Lancaster"
    },
    "company": "HANDSHAKE",
    "email": "wall.lancaster@handshake.ca",
    "phone": "+1 (972) 420-3407",
    "address": "108 Ferry Place, Epworth, West Virginia, 5970",
    "about": "Mollit aliqua officia cillum Lorem eu ipsum ipsum labore anim proident qui pariatur enim. Mollit dolor laboris esse consectetur pariatur consectetur proident non nostrud eiusmod elit. In occaecat excepteur nisi ad excepteur nostrud incididunt sunt anim occaecat sit magna quis. Sint sunt excepteur mollit excepteur occaecat ullamco eiusmod dolor cillum tempor et Lorem. Sint commodo laboris nisi laborum adipisicing dolore in exercitation. Do nostrud minim non exercitation. Duis quis ullamco voluptate pariatur sit elit deserunt dolore.",
    "registered": "Sunday, June 16, 2019 9:59 AM",
    "latitude": "87.860956",
    "longitude": "74.936819",
    "tags": [
      "eiusmod",
      "do",
      "mollit",
      "laborum",
      "pariatur"
    ],
    "range": [
      0,
      1,
      2,
      3,
      4,
      5,
      6,
      7,
      8,
      9
    ],
    "friends": [
      {
        "id": 0,
        "name": "Muriel Hood"
      },
      {
        "id": 1,
        "name": "Ashley Trujillo"
      },
      {
        "id": 2,
        "name": "Davis Collins"
      }
    ],
    "greeting": "Hello, Wall! You have 9 unread messages.",
    "favoriteFruit": "strawberry"
  },
  {
    "_id": "5f5830a17646558f8743dc73",
    "index": 3,
    "guid": "c7c30f87-600e-451e-9ac7-daa02251080e",
    "isActive": false,
    "balance": "$1,530.45",
    "picture": "http://placehold.it/32x32",
    "age": 32,
    "eyeColor": "brown",
    "name": {
      "first": "Flowers",
      "last": "Abbott"
    },
    "company": "COMTRACT",
    "email": "flowers.abbott@comtract.co.uk",
    "phone": "+1 (872) 486-2612",
    "address": "673 Bergen Street, Wintersburg, New Mexico, 1994",
    "about": "Eiusmod cillum eu aute eiusmod ipsum ullamco. Duis voluptate consectetur culpa in in et dolore irure sit laboris aliqua et. Nulla quis ea anim eiusmod non et pariatur laboris officia. Aute magna culpa magna eu sint et nisi.",
    "registered": "Thursday, January 23, 2014 6:12 PM",
    "latitude": "78.700956",
    "longitude": "120.752362",
    "tags": [
      "eiusmod",
      "non",
      "commodo",
      "reprehenderit",
      "id"
    ],
    "range": [
      0,
      1,
      2,
      3,
      4,
      5,
      6,
      7,
      8,
      9
    ],
    "friends": [
      {
        "id": 2,
        "name": "Morse Le"
      },
      {
        "id": 1,
        "name": "Magdalena Reyes"
      },
      {
        "id": 2,
        "name": "Maldonado Gray"
      }
    ],
    "greeting": "Hello, Flowers! You have 7 unread messages.",
    "favoriteFruit": "apple"
  },
  {
    "_id": "5f5830a1e12f937da6a1c728",
    "index": 4,
    "guid": "9f1a9248-97f6-4e75-8fda-668da6694059",
    "isActive": true,
    "balance": "$3,492.33",
    "picture": "http://placehold.it/32x32",
    "age": 37,
    "eyeColor": "brown",
    "name": {
      "first": "Blevins",
      "last": "Woods"
    },
    "company": "RUGSTARS",
    "email": "blevins.woods@rugstars.io",
    "phone": "+1 (870) 488-3919",
    "address": "762 Montgomery Place, Sardis, Tennessee, 4800",
    "about": "Esse ut aute duis mollit. Esse tempor eiusmod nostrud sint amet fugiat fugiat culpa ad incididunt ad. Cillum incididunt incididunt laboris non pariatur amet et adipisicing adipisicing fugiat ut cillum. Reprehenderit dolor ullamco occaecat ipsum qui et. Minim sint non amet officia culpa. Eu eiusmod Lorem eu aute cillum commodo aliquip aliquip qui. Elit Lorem adipisicing exercitation nisi anim sit deserunt culpa proident consequat cupidatat cupidatat amet pariatur.",
    "registered": "Monday, September 30, 2019 7:11 PM",
    "latitude": "-60.846886",
    "longitude": "-78.994875",
    "tags": [
      "duis",
      "veniam",
      "incididunt",
      "proident",
      "proident"
    ],
    "range": [
      0,
      1,
      2,
      3,
      4,
      5,
      6,
      7,
      8,
      9
    ],
    "friends": [
      {
        "id": 0,
        "name": "Morgan Stanton"
      },
      {
        "id": 1,
        "name": "Pearson Fischer"
      },
      {
        "id": 2,
        "name": "Roach Skinner"
      }
    ],
    "greeting": "Hello, Blevins! You have 5 unread messages.",
    "favoriteFruit": "apple"
  },
  {
    "_id": "5f5830a117fbe9268e20244f",
    "index": 5,
    "guid": "00f6348c-3cc8-4b6e-b7ad-6ac425d0cb19",
    "isActive": false,
    "balance": "$3,809.17",
    "picture": "http://placehold.it/32x32",
    "age": 24,
    "eyeColor": "blue",
    "name": {
      "first": "Angelica",
      "last": "Maldonado"
    },
    "company": "MEDESIGN",
    "email": "angelica.maldonado@medesign.info",
    "phone": "+1 (936) 565-3922",
    "address": "499 Kensington Walk, Grenelefe, Delaware, 2200",
    "about": "Occaecat ipsum labore velit mollit sit labore ipsum. Commodo Lorem dolor incididunt eu sit aliquip minim exercitation eiusmod laboris commodo quis do. Consectetur commodo veniam consectetur ipsum duis aliqua non consequat sint officia enim id. Exercitation voluptate adipisicing fugiat labore officia occaecat ullamco fugiat ex. Lorem occaecat est do sunt pariatur veniam.",
    "registered": "Saturday, May 9, 2015 2:27 PM",
    "latitude": "34.07016",
    "longitude": "-96.727205",
    "tags": [
      "sit",
      "adipisicing",
      "aliquip",
      "id",
      "culpa"
    ],
    "range": [
      0,
      1,
      2,
      3,
      4,
      5,
      6,
      7,
      8,
      9
    ],
    "friends": [
      {
        "id": 0,
        "name": "Keller House"
      },
      {
        "id": 1,
        "name": "King Soto"
      },
      {
        "id": 2,
        "name": "Salas Alford"
      }
    ],
    "greeting": "Hello, Angelica! You have 7 unread messages.",
    "favoriteFruit": "banana"
  }
]`

func GetJSONData() interface{} {
	var d interface{}

	err := json.Unmarshal([]byte(JSONData), &d)
	if err != nil {
		panic(err)
	}

	return d
}

func GetJSONData2() interface{} {
	var d interface{}

	err := json.Unmarshal([]byte(JSONData2), &d)
	if err != nil {
		panic(err)
	}

	return d
}

