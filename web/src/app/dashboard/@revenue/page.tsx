"use client"
import Card from "@/src/app/ui/card";
import {useEffect, useState} from "react";

export default function RevenueMetrics() {
    const [data, setData] = useState();
    const [loading, setLoading] = useState(true);  // Start with loading true
    const [error, setError] = useState("");


    useEffect(() => {
        setTimeout(() => {
            try {
                const fetchedData = {id: 1, name: 'John Doe'};
                setData(fetchedData as any);
                setLoading(false);
            } catch (e) {
                setError('Failed to fetch data');
                setLoading(false);
            }
        }, 2000)
    }, [])


    if (error) {
        return <div>Error: {error}</div>;
    }

    return <Card>
        {loading ? <div>RevenueMetrics Loading...</div> : <div>RevenueMetrics</div>}
    </Card>
}
