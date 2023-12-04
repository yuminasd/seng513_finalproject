export interface Group {
    id: string;
    name: string;
    members: any[];
    likedMovies: any[];
}

export interface User {
    id: string;
    name: string;
    img: string;
    password?: string;
    // movies: Movie[];
}

export interface Movie {
    id: string;
    name: string;
    img: string;
    bgImg: string;
    rating: number;
    description: string;
    genres: [];

}
