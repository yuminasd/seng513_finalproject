'use client'
import Table from "@/app/components/table";
import Navbar from "../../components/navbar";
import MovieCard from "@/app/components/movieCard";
import Button from "@/app/components/button";
import { groupsMock } from "@/app/mock";

export default function Home({ params }: { params: { slug: number } }) {
    const columns = ['name'];
    return (
        <section className="flex p-4 ">
            <Navbar />
            {/* Users List */}
            <div className="w-5/12 px-4 flex flex-col gap-4 border-r border-neutral-900 max-sm:hidden">
                <div className="w-full flex flex-col justify-center text-center gap-2 p-16 border border-neutral-900 rounded-2xl">
                    <h1 className="text-2xl font-bol">{groupsMock[params.slug - 1].groupName}</h1>
                    <p> {params.slug}</p>
                </div>
                <Table columns={columns} data={groupsMock[params.slug - 1].users} page="group" />

            </div>

            {/* Movies List */}
            <div className=" grid lg:grid-cols-4 md:grid-cols-3 sm:grid-cols-1  gap-4 justify-start h-screen overflow-y-auto content-start px-4 ">
                {groupsMock[params.slug - 1].movies.map((movie, index) => (
                    <MovieCard key={index} movie={movie} />
                ))}
            </div>
            <div className="absolute left-0 bottom-0 w-full bg-neutral-900 p-4">
                <Button text={"Swipe"} color={"primary"} />
            </div>
        </section>
    )
}
