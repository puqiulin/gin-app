import ClientComponent from "@/src/app/ui/client-component";
import ServerComponent from "@/src/app/ui/server-component";

export default function Interleaving() {
    return <>
        <h1>Interleaving</h1>

        {/*this will pop up error*/}
        {/*<ClientComponent/>*/}
        {/*<ServerComponent/>*/}

        <ClientComponent>
            <ServerComponent/>
        </ClientComponent>
    </>
}
