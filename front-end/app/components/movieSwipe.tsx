import React, { useState } from 'react';
import { Movie } from '../types';

const MovieSwipe: React.FC<{ movie: Movie }> = ({ movie }) => {


    return (
        <div className="overflow-hidden relative flex flex-col gap-1 bg-black rounded-md  w-full h-full aspect-video border border-neutral-900 p-4">
            <p className="text-center">{movie.name}</p>
            <img className="object-contain h-[20rem] " src={movie.bgImg} alt={`${movie.name} background`} />
            <p>Rating: {movie.rating}</p>
            <p className="w-full"> {movie.description}</p>
        </div>
    );
};

export default MovieSwipe;
