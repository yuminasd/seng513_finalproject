'use client';
import { useState, useEffect } from 'react'
import Navbar from "../components/navbar";
import { Movie } from "../types";
import MovieCard from "../components/movieCard";
import { moviesMock } from '../mock';
import MovieSwipe from '../components/movieSwipe';


function Swipe() {

    const [movie1, setMovie1] = useState<Movie>(moviesMock[0]);
    const [movie2, setMovie2] = useState<Movie>(moviesMock[1]);
    const [movie3, setMovie3] = useState<Movie>(moviesMock[2]);

    const fetchRandomMovie = async (usedIndexes) => {
        try {
            const response = await fetch('http://localhost:5000/movies');
            const data = await response.json();

            // Create an array of all indexes
            const allIndexes = Array.from({ length: data.data.data.length }, (_, index) => index);

            // Exclude the usedIndexes from allIndexes
            const availableIndexes = allIndexes.filter(index => !usedIndexes.includes(index));

            // If there are available indexes, select a random one
            if (availableIndexes.length > 0) {
                const randomIndex = availableIndexes[Math.floor(Math.random() * availableIndexes.length)];
                const randomMovie = data.data.data[randomIndex];
                return { randomMovie, randomIndex };
            } else {
                // All indexes have been used, handle this case based on your requirements
                console.warn('No available indexes. All movies have been used.');
                return null;
            }
        } catch (error) {
            console.error('Error fetching movies data:', error);
            return null;
        }
    };



    const fetchData = async () => {
        try {
            const { randomMovie: movie1, randomIndex: index1 } = await fetchRandomMovie([]);
            const { randomMovie: movie2, randomIndex: index2 } = await fetchRandomMovie([index1]);
            const { randomMovie: movie3, randomIndex: index3 } = await fetchRandomMovie([index1, index2]);

            setMovie1(movie1);
            setMovie2(movie2);
            setMovie3(movie3);
        } catch (error) {
            console.error('Error fetching movies data:', error);
        }
    };
    useEffect(() => {
        fetchData();
    }, []);

    const handleLike = async () => {
        try {

            const userId = "656500413c49f1af1a59b5d1";
            const movieId = movie2.id; // Assuming you want to like the first movie
            try {

                const response = await fetch(`http://localhost:5000/addliked/${userId}/${movieId}`, {
                    method: 'PATCH',
                    headers: {
                        'Content-Type': 'application/json',
                    }
                });

                const data = await response.json();
                console.log(data);
            } catch (error) {
                console.error('Error fetching adding liked movie', error);
            }
            fetchData();
        } catch (error) {
            console.error('Error liking movie:', error);
        }
    };
    return (
        <section className="flex flex-col h-screen overflow-hidden">
            <Navbar />
            <div className="flex h-full">
                <div className="h-96 w-[1/3] flex-col flex justify-end items-center rounded-xl overflow-hidden p-4 brightness-50 max-md:hidden">
                    Disliked Movies
                    <img className="object-fill" src={movie1.bgImg} />
                </div>
                <MovieSwipe movie={movie2} />
                <div className="h-96 w-[1/3] flex-col flex justify-end items-center rounded-xl overflow-hidden p-4 brightness-50 max-md:hidden">
                    Liked Movies
                    <img className="object-fill" src={movie3.bgImg} />
                </div>
            </div>
            <div className="absolute bg-black border border-t border-neutral-800 bottom-0 left-0 w-full flex p-4 gap-4 ">
                <button className=" rounded bg-red-500  py-2 w-full" onClick={fetchData}> Dislike</button>
                <button className="rounded bg-green-500 py-2 w-full" onClick={handleLike}> Like</button>
            </div>
        </section>
    );
};

export default Swipe