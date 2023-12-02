import { Group, Movie, User } from "./types"

export let moviesMock: Movie[] = [{
    name: "The OA",
    img: "https://occ-0-1168-299.1.nflxso.net/dnm/api/v6/6gmvu2hxdfnQ55LZZjyzYR4kzGk/AAAABVXuj1XWKq_TF-WSc6Pq0Dl5VSHmkOiCMHIUzZ3AyZg7NWd2JoFfeQvHNBdOJGKct_dhYAbC_3DGNAG-oHxaMzBnXB7AnEjqc0knZLgTgeo3e-nYN0CLRHrMPjC5YppyUup1.jpg?r=85c",
    rating: 500,
},
{
    name: "Love, Death & Robots",
    img: "https://occ-0-1168-299.1.nflxso.net/dnm/api/v6/6gmvu2hxdfnQ55LZZjyzYR4kzGk/AAAABZJ_7CPwddfqeyjXjyyCCk_UqWvDv04NbS4g5GDfBOYlynmWzTgSHuRWCbB63Y3tLBILZ5mzWD7DGNSTGhJfBpkq4-t_bLeZzTHqAd5ROz6SNWm7hLGFlPTxaJKgwWTB7oWS.jpg?r=73",
    rating: 500,
},

{
    name: "DARK",
    img: "https://occ-0-1168-299.1.nflxso.net/dnm/api/v6/6gmvu2hxdfnQ55LZZjyzYR4kzGk/AAAABYuKjjKMzFh6RN2q7ml5nJDrNguQzImqCdD7tTYKKUXYnliwYVjasxzPERdiwpcDhL8zqUyaRHnRhYWtajxpwaRLf1FUTtHR7CGGFkn028rZF0CCpGfBAYA-e6H0DdaVUrdO.jpg?r=393",
    rating: 500,
},
{
    name: "Altered Carbon",
    img: " https://occ-0-1168-299.1.nflxso.net/dnm/api/v6/6gmvu2hxdfnQ55LZZjyzYR4kzGk/AAAABdiUookKGsY-jFH8p-YKtpeSZ0JD1UeGrzMd946ERVCw16h6_msxXrWa3Oi3lc3Hggl0ogaZ7BqDZslDSqJIPAd4860et1KFT9AdYGbUQmOLZy08I1Od6RasNF3hc1Oe8pS_.jpg?r=525", rating: 500,
},
{
    name: "Travelers",
    img: "https://occ-0-1168-299.1.nflxso.net/dnm/api/v6/6gmvu2hxdfnQ55LZZjyzYR4kzGk/AAAABd9eAePC0rHO92Jv8Yo-vlBQ6vylDNRKqBI63XNZN4RmdL9LmyuWlBfXMGL5bEn90TT1uUmK0tawHJ_a9cbhrtuPjn3oDcifFhyDejf0Ksb56nLqqVSpnWtOwhcKg0YgJ9zH.jpg?r=b93", rating: 500,
},]

export let usersMock: User[] =
    [{
        name: "John Doe",
        img: "",
    },
    {
        name: "Big BOOOOIII",
        img: "",
    },]



export let groupsMock: Group[] = [{
    groupName: "Group1",
    code: "1",
    users: usersMock,
    movies: moviesMock,
},
{
    groupName: "Group2",
    code: "2",
    users: usersMock,
    movies: moviesMock,
},
]