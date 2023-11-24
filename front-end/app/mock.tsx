import { Group, Movie, User } from "./types"

export let moviesMock: Movie[] = [{
    name: "Stranger Things",
    img: "https://occ-0-1168-299.1.nflxso.net/dnm/api/v6/E8vDc_W8CLv7-yMQu8KMEC7Rrr8/AAAABSpqYf-qwjTo6Bv_bRrj9_-f_Pi-CYcNZ1ICnyrp7a8CD2Cu4TyrntigqIw7CMFMj8f_Lr3zHq6G2WtkHs_bc8vsqISuGHe4lFyT.webp?r=fc0",
    rating: 500,
},
{
    name: "Love, Death & Robots",
    img: "https://occ-0-1168-299.1.nflxso.net/dnm/api/v6/6gmvu2hxdfnQ55LZZjyzYR4kzGk/AAAABZJ_7CPwddfqeyjXjyyCCk_UqWvDv04NbS4g5GDfBOYlynmWzTgSHuRWCbB63Y3tLBILZ5mzWD7DGNSTGhJfBpkq4-t_bLeZzTHqAd5ROz6SNWm7hLGFlPTxaJKgwWTB7oWS.jpg?r=73",
    rating: 500,
},

{
    name: "Stranger Things",
    img: "https://occ-0-1168-299.1.nflxso.net/dnm/api/v6/6gmvu2hxdfnQ55LZZjyzYR4kzGk/AAAABYuKjjKMzFh6RN2q7ml5nJDrNguQzImqCdD7tTYKKUXYnliwYVjasxzPERdiwpcDhL8zqUyaRHnRhYWtajxpwaRLf1FUTtHR7CGGFkn028rZF0CCpGfBAYA-e6H0DdaVUrdO.jpg?r=393",
    rating: 500,
},
{
    name: "Stranger Things",
    img: "https://occ-0-1168-299.1.nflxso.net/dnm/api/v6/E8vDc_W8CLv7-yMQu8KMEC7Rrr8/AAAABSpqYf-qwjTo6Bv_bRrj9_-f_Pi-CYcNZ1ICnyrp7a8CD2Cu4TyrntigqIw7CMFMj8f_Lr3zHq6G2WtkHs_bc8vsqISuGHe4lFyT.webp?r=fc0",
    rating: 500,
},
{
    name: "Stranger Things",
    img: "https://occ-0-1168-299.1.nflxso.net/dnm/api/v6/E8vDc_W8CLv7-yMQu8KMEC7Rrr8/AAAABSpqYf-qwjTo6Bv_bRrj9_-f_Pi-CYcNZ1ICnyrp7a8CD2Cu4TyrntigqIw7CMFMj8f_Lr3zHq6G2WtkHs_bc8vsqISuGHe4lFyT.webp?r=fc0",
    rating: 500,
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