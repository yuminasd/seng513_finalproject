'use client';
import { User } from "@/app/types";
import Navbar from "../../components/navbar";
import { useState, useEffect } from 'react'
import MovieCard from "@/app/components/movieCard";

export default function Profile({ params }: { params: { slug: number } }) {
    const [user, setUser] = useState<User | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch('http://localhost:5000/users/' + params.slug, {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                    }
                });
                const data = await response.json();
                console.log(data.data.data);
                setUser(data.data.data as User);
            } catch (error) {
                console.error('Error fetching groups data:', error);
            }
        };


        fetchData();
    }, []);
    return (
        <section className=" flex p-4 h-screen">
            <Navbar />
            <form className="flex flex-col gap-4 p-4  border-r border-neutral-800 h-full">
                <img className="p-4 h-36  bg-cover border border-neutral-800 rounded" src={user?.img} alt="User Image" />
                <input className="bg-black p-2  border border-neutral-800 rounded " readOnly defaultValue={user?.name} />
                <input className="bg-black p-2  border border-neutral-800 rounded " readOnly defaultValue={user?.id} />

            </form>
            <div className=" flex flex-col gap-4  text-center ">
                <p> List of all liked movies</p>
                {/* {user?.likedMovies ? (
                    <div className=" grid lg:grid-cols-4 md:grid-cols-3 sm:grid-cols-1  gap-4 justify-start h-screen overflow-y-auto content-start px-4 ">
                        {user.likedMovies.map((movie, index) => (
                            <MovieCard key={index} movie={movie.movieId} />
                        ))}
                    </div>
                ) : (
                    <p>Loading Movies data...</p>
                )
                } */}
            </div>

        </section>
    );
};
