'use client'

import Button from "./button";

export default function Swipe() {
    return (
        <div className="absolute left-0 bottom-0 w-full bg-neutral-900 p-4">
            <Button text={"Swipe"} color={"primary"} />
        </div>
    )
}