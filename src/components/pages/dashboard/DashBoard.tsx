import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Button } from "@mui/material";
import { logOut } from "../../utils/HelperFuncs";
import { familyList } from "../../utils/JsonExample";
import { DataGrid, GridColDef, GridRowsProp } from '@mui/x-data-grid';

export const DashBoard = () => {
    const rows: GridRowsProp = [
        { id: 1, col1: 'Hello', col2: 'World' },
        { id: 2, col1: 'DataGridPro', col2: 'is Awesome' },
        { id: 3, col1: 'MUI', col2: 'is Amazing' },
    ];

    const columns: GridColDef[] = [
        { field: 'col1', headerName: 'Column 1', width: 150 },
        { field: 'col2', headerName: 'Column 2', width: 150 },
    ];

    return (
        <div style={{ height: 300, width: '100%' }}>
            <DataGrid rows={rows} columns={columns} />
            <svg width="100vw" height="500" >
                <circle cx="50vw" cy="250" r="50" fill="yellow" stroke="black" stroke-width="2"></circle>
            </svg>
            <Button variant="contained">Hi</Button>
        </div>
    );
}
