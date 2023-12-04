'use client'
import Table from "@/app/components/table";
import Navbar from "../../components/navbar";
import MovieCard from "@/app/components/movieCard";
import Button from "@/app/components/button";
import { useRouter } from 'next/navigation'
import { useState, useEffect } from 'react'
import { Group, Movie } from "@/app/types";
import Swipe from "@/app/components/swipe";


export default function Group({ params }: { params: { slug: number } }) {
    const columns = ['name', 'id'];
    const [group, setGroup] = useState<Group | null>(null);

    // const [movies, setMovies] = useState<Movie[] | []>([]);
    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch('http://localhost:5000/groups/' + params.slug, {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                    }
                });
                const data = await response.json();
                console.log(data.data.data);
                setGroup(data.data.data as Group);

                // data.data.data.likedMovies.forEach(function(movie: {likedCount:string,movieId:string }) {
                //     await fetch("http://localhost:5000/movies")
                // })

                // const moviesResponse = await fetch('http://localhost:5000/movies', {
                //     method: 'GET',
                //     headers: {
                //         'Content-Type': 'application/json',
                //     }
                // });
                // const movieResponse = await moviesResponse.json();
                // console.log(movieResponse.data.data);
                // setMovies(movieResponse.data.data);
            } catch (error) {
                console.error('Error fetching groups data:', error);
            }
        };


        fetchData();
    }, []);


    return (

        <section className="flex p-4 ">
            <Navbar />
            {/* Users List */}
            <div className="w-5/12 px-4 flex flex-col gap-4 border-r border-neutral-900 max-sm:hidden">
                <div className="w-full flex flex-col justify-center text-center gap-2 p-16 border border-neutral-900 rounded-2xl">
                    {group ? (
                        <div>
                            <h1 className="text-2xl font-bol">{group?.name}</h1>
                            <p> {group?.id}</p>
                        </div>
                    ) : (
                        <p>Loading group data...</p>
                    )}
                </div>

                {group ? (
                    <Table columns={columns} data={group?.members} page="group" />
                ) : (
                    <p>Loading group data...</p>
                )}
            </div>

            {group?.likedMovies ? (

                <div className=" grid lg:grid-cols-4 md:grid-cols-3 sm:grid-cols-1  gap-4 justify-start h-screen overflow-y-auto content-start px-4 ">
                    {group.likedMovies.map((movie, index) => (
                        <MovieCard key={index} movie={movie.movieId} />
                    ))}
                </div>
            ) : (
                <p>Loading Movies data...</p>
            )
            }
            <Swipe />

        </section>
    )
}
