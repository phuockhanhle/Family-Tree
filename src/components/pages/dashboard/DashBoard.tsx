import { Button } from "@mui/material";
import * as d3 from "d3";
import { useCallback, useEffect, useState } from "react";

const width = '100vw';
const height = 300;
const circleX = '50vw';
const circleY = height / 2;
const circleRadius = 50;
const initialMousePosition = {
    x: circleX,
    y: height / 2
}
const csvUrl = 'https://gist.githubusercontent.com/curran/b236990081a24761f7000567094914e0/raw/cssNamedColors.csv'

export const DashBoard = () => {
    const [mousePosition, setMousePosition] = useState(initialMousePosition);
    const [message, setMessage] = useState('')
    const [data, setData] = useState(null);

    // wont create a new version of 'handleMouseMove' if setMousePosition doesnt change
    const handleMouseMove = useCallback((event: any) => {
        const { clientX, clientY } = event;
        setMousePosition({ x: clientX, y: clientY });
    }, [setMousePosition])

    useEffect(() => {
        d3.csv(csvUrl).then(data => {
            let message = '';
            message = message + Math.round(d3.csvFormat(data).length / 1024) + ' kB\n';
            message = message + data.length + ' rows\n';
            message = message + data.columns.length + ' columns';
            setMessage(message);
        })
    }, [])

    

    return (
        <>
            <svg width={width} height={height} onMouseMove={handleMouseMove}>
                <g>
                    <circle cx={mousePosition.x} cy={mousePosition.y} r={circleRadius} fill="yellow" stroke="black" strokeWidth="2"></circle>
                </g>
            </svg>
            <Button variant="contained">{message}</Button>
        </>


    );
}
