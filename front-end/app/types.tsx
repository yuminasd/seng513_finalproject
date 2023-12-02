export interface Group {
    groupName: string;
    code: string;
    users: User[];
    movies: Movie[];
}

export interface User {
    name: string;
    img: string;
    password?: string;
    // movies: Movie[];
}

export interface Movie {
    name: string;
    img: string;
    rating: number;
}
