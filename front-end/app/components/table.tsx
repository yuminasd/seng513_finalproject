'use client'
import { useRouter } from 'next/navigation'
import React from 'react';

// Define the props interface for the DynamicTable component
interface TableProps {
    columns: string[]; // Array of column names
    data: Record<string, any>[]; // Array of data objects, each representing a row
    page: string;

}

const Table: React.FC<TableProps> = ({ columns, data, page }) => {
    const router = useRouter();
    const navigate = (groupCode: string) => {
        if (page === "groups") {
            router.push(`/group/${groupCode}`);
        }
        else if (page === "group") {
            router.push(`/profile`);
        }

    }

    return (
        <table className="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
            <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-opacity-10 dark:text-gray-400">
                <tr>
                    {/* Map through columns to create table header cells */}
                    {columns.map((column, index) => (
                        <th className="px-6 py-3" key={index}>{column}</th>
                    ))}
                </tr>
            </thead>

            <tbody>
                {/* Map through data to create table rows */}
                {data.map((row, rowIndex) => (
                    <tr className="bg-white border-b dark:bg-opacity-5 dark:border-white dark:border-opacity-10 hover:bg-opacity-10" key={rowIndex}>
                        {/* Map through columns to create table cells for each row */}
                        {columns.map((column, colIndex) => (
                            <td className="px-6 py-4" key={colIndex} onClick={() => navigate(row[columns[1]])}>{row[column]}</td>
                        ))}
                    </tr>
                ))}
            </tbody>
        </table>
    );
};

export default Table;