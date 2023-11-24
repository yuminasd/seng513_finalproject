'use client'
import Table from "@/app/components/table";
import Navbar from "../../components/navbar";
import MovieCard from "@/app/components/movieCard";
import Button from "@/app/components/button";
import { groupsMock } from "@/app/mock";

export default function Home({ params }: { params: { slug: string } }) {
    const columns = ['name'];
    return (
        <main className="flex min-h-screen min-w-screen">
            <Navbar />
            <div className="w-5/12 pt-24 p-4 flex flex-col gap-4 ">
                <div className="w-full flex flex-col justify-center text-center gap-2 p-16 border border-gray-700 rounded-2xl">
                    <h1 className="text-2xl font-bol">{groupsMock[0].groupName}</h1>
                    <p> {params.slug}</p>
                </div>
                <Table columns={columns} data={groupsMock[0].users} />

            </div>


            <div className="pt-24 flex flex-wrap gap-4 justify-start h-screen overflow-y-auto content-start ">

                {groupsMock[0].movies.map((movie, index) => (
                    // <div className=" overflow-hidden relative flex flex-col w-[20rem] h-[18rem] bg-white bg-opacity-10 rounded-md ">
                    //     <img className="absolute bottom-8 self-center w-64 " src={movie.img} />
                    //     <img className="object-cover w-[25rem] h-[25rem] " src={movie.bgImg} />
                    //     {/* <p key={index} className="text-center absolute bottom-4 w-full">
                    //         {movie.name}
                    //     </p> */}
                    // </div>
                    <MovieCard key={index} movie={movie} />

                ))}
            </div>
            <div className="absolute bottom-0 w-full bg-neutral-900 p-4">
                <Button text={"Swipe"} color={"primary"} />
            </div>
        </main>
    )
}
