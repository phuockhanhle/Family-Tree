import { Button } from "@mui/material";
import { useState } from "react";

const width = '100vw';
const height = 300;
const circleX = '50vw';
const circleY = height / 2;
const circleRadius = 50;
const initialMousePosition = {
    x: circleX,
    y: height / 2
}


export const DashBoard = () => {
    const [mousePosition, setMousePosition] = useState(initialMousePosition);
    const handleMouseMove = (event: any) => {
        const { clientX, clientY } = event;
        console.log({ clientX, clientY });
        setMousePosition({ x: clientX, y: clientY });
    }

    return (
        <>
            <svg width={width} height={height} onMouseMove={handleMouseMove}>
                <g>
                    <circle cx={mousePosition.x} cy={mousePosition.y} r={circleRadius} fill="yellow" stroke="black" stroke-width="2"></circle>
                </g>
            </svg>
            <Button variant="contained">Hi</Button>
        </>


    );
}
