"use client"

import {gql, useQuery} from '@apollo/client';
import {Key} from 'react';

export default function Graphql() {
    const GET_USERS = gql`
        query {
            users{
                id
                name
                email
                gender
                phone
                last_login_time
                last_login_ip
                create_time
                update_time
                posts{
                    id
                    title
                }
            }
        }
    `;

    const {loading, error, data} = useQuery(GET_USERS);

    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error : {error.message}</p>;

    return <div>
        {data.users.map((u: any, i: Key) => <div key={i} className="flex gap-2">
                <div>{u.name}</div>
                <div>{u.email}</div>
                <div>{u.phone}</div>
                <div>{u.create_time}</div>
            </div>
        )}
    </div>
}