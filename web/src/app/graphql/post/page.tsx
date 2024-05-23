"use client"
import {useQuery, gql} from '@apollo/client';
import {Key} from 'react';

export default function Graphql() {
    const GET_USERS = gql`
        query {
            posts{
                id
                title
                user{
                    id
                    name
                }
            }
        }
    `;

    const {loading, error, data} = useQuery(GET_USERS);

    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error : {error.message}</p>;

    return <div>
        {data.posts.map((p: any, i: Key) => <div key={i} className="flex gap-2">
                <div>{p.title}</div>
                <div>{p.user.name}</div>
                <div>{p.create_time}</div>
            </div>
        )}
    </div>
}