'use client'
import Table from "@/app/components/table";
import Navbar from "../../components/navbar";
import MovieCard from "@/app/components/movieCard";
import Button from "@/app/components/button";

export default function Home({ params }: { params: { slug: string } }) {


    let moviesMock = [{
        name: "Stranger Things",
        img: "https://occ-0-1168-299.1.nflxso.net/dnm/api/v6/tx1O544a9T7n8Z_G12qaboulQQE/AAAABSNUzoP8HHEMoUSYEp5GdQ92BwmUzjENlcVRytYB-9zH3CWjN1d3IEkLGjB8njlIALYjHH8NG9eT0k876YHi9_JtUIJCaG9ZXHcEC26sYbfa-GlUN6Xyuvg5dqpkUDPvoUvTLYssGgkv0YErnLL1LEzancT6V39BpxajPFO7uP1Q_SpE-D4E.webp?r=23b",
        bgImg: "https://occ-0-1168-299.1.nflxso.net/dnm/api/v6/E8vDc_W8CLv7-yMQu8KMEC7Rrr8/AAAABSpqYf-qwjTo6Bv_bRrj9_-f_Pi-CYcNZ1ICnyrp7a8CD2Cu4TyrntigqIw7CMFMj8f_Lr3zHq6G2WtkHs_bc8vsqISuGHe4lFyT.webp?r=fc0",
        rating: 500,
    },
    {
        name: "Love, Death & Robots",
        img: "https://occ-0-1168-299.1.nflxso.net/dnm/api/v6/tx1O544a9T7n8Z_G12qaboulQQE/AAAABSNUzoP8HHEMoUSYEp5GdQ92BwmUzjENlcVRytYB-9zH3CWjN1d3IEkLGjB8njlIALYjHH8NG9eT0k876YHi9_JtUIJCaG9ZXHcEC26sYbfa-GlUN6Xyuvg5dqpkUDPvoUvTLYssGgkv0YErnLL1LEzancT6V39BpxajPFO7uP1Q_SpE-D4E.webp?r=23b",
        bgImg: "https://occ-0-1168-299.1.nflxso.net/dnm/api/v6/6gmvu2hxdfnQ55LZZjyzYR4kzGk/AAAABZJ_7CPwddfqeyjXjyyCCk_UqWvDv04NbS4g5GDfBOYlynmWzTgSHuRWCbB63Y3tLBILZ5mzWD7DGNSTGhJfBpkq4-t_bLeZzTHqAd5ROz6SNWm7hLGFlPTxaJKgwWTB7oWS.jpg?r=73",
        rating: 500,
    },

    {
        name: "Stranger Things",
        img: "https://occ-0-1168-299.1.nflxso.net/dnm/api/v6/tx1O544a9T7n8Z_G12qaboulQQE/AAAABSNUzoP8HHEMoUSYEp5GdQ92BwmUzjENlcVRytYB-9zH3CWjN1d3IEkLGjB8njlIALYjHH8NG9eT0k876YHi9_JtUIJCaG9ZXHcEC26sYbfa-GlUN6Xyuvg5dqpkUDPvoUvTLYssGgkv0YErnLL1LEzancT6V39BpxajPFO7uP1Q_SpE-D4E.webp?r=23b",
        bgImg: "https://occ-0-1168-299.1.nflxso.net/dnm/api/v6/6gmvu2hxdfnQ55LZZjyzYR4kzGk/AAAABYuKjjKMzFh6RN2q7ml5nJDrNguQzImqCdD7tTYKKUXYnliwYVjasxzPERdiwpcDhL8zqUyaRHnRhYWtajxpwaRLf1FUTtHR7CGGFkn028rZF0CCpGfBAYA-e6H0DdaVUrdO.jpg?r=393",
        rating: 500,
    },
    {
        name: "Stranger Things",
        img: "https://occ-0-1168-299.1.nflxso.net/dnm/api/v6/tx1O544a9T7n8Z_G12qaboulQQE/AAAABSNUzoP8HHEMoUSYEp5GdQ92BwmUzjENlcVRytYB-9zH3CWjN1d3IEkLGjB8njlIALYjHH8NG9eT0k876YHi9_JtUIJCaG9ZXHcEC26sYbfa-GlUN6Xyuvg5dqpkUDPvoUvTLYssGgkv0YErnLL1LEzancT6V39BpxajPFO7uP1Q_SpE-D4E.webp?r=23b",
        bgImg: "https://occ-0-1168-299.1.nflxso.net/dnm/api/v6/E8vDc_W8CLv7-yMQu8KMEC7Rrr8/AAAABSpqYf-qwjTo6Bv_bRrj9_-f_Pi-CYcNZ1ICnyrp7a8CD2Cu4TyrntigqIw7CMFMj8f_Lr3zHq6G2WtkHs_bc8vsqISuGHe4lFyT.webp?r=fc0",
        rating: 500,
    },
    {
        name: "Stranger Things",
        img: "https://occ-0-1168-299.1.nflxso.net/dnm/api/v6/tx1O544a9T7n8Z_G12qaboulQQE/AAAABSNUzoP8HHEMoUSYEp5GdQ92BwmUzjENlcVRytYB-9zH3CWjN1d3IEkLGjB8njlIALYjHH8NG9eT0k876YHi9_JtUIJCaG9ZXHcEC26sYbfa-GlUN6Xyuvg5dqpkUDPvoUvTLYssGgkv0YErnLL1LEzancT6V39BpxajPFO7uP1Q_SpE-D4E.webp?r=23b",
        bgImg: "https://occ-0-1168-299.1.nflxso.net/dnm/api/v6/E8vDc_W8CLv7-yMQu8KMEC7Rrr8/AAAABSpqYf-qwjTo6Bv_bRrj9_-f_Pi-CYcNZ1ICnyrp7a8CD2Cu4TyrntigqIw7CMFMj8f_Lr3zHq6G2WtkHs_bc8vsqISuGHe4lFyT.webp?r=fc0",
        rating: 500,
    },]

    let groupMock = {
        groupName: "bob",
        code: "123",
        users: [{
            username: "123",
            img: "",
        },
        {
            username: "123",
            img: "",
        },
        ],
        movies: moviesMock,
    };

    const columns = ['username'];
    return (
        <main className="flex min-h-screen min-w-screen">
            <Navbar />
            <div className="w-5/12 pt-24 p-4 flex flex-col gap-4 ">
                <div className="w-full flex flex-col justify-center text-center gap-2 p-16 border border-gray-700 rounded-2xl">
                    <h1 className="text-2xl font-bol">{groupMock.groupName}</h1>
                    <p> {params.slug}</p>
                </div>
                <Table columns={columns} data={groupMock.users} />

            </div>


            <div className="pt-24 flex flex-wrap gap-4 justify-start h-screen overflow-y-auto content-start ">

                {moviesMock.map((movie, index) => (
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
