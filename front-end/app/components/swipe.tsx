'use client'
import { useRouter } from "next/navigation";
import Button from "./button";


export default function Swipe() {
    const router = useRouter();

    const navigate = () => {
        router.push(`/swipe`);
    }

    return (
        <div className="absolute left-0 bottom-0 w-full bg-neutral-900 p-4">
            <Button text={"Swipe"} color={"primary"} onClick={navigate} />
        </div>
    )
}