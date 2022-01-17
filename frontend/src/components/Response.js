export const Response = ({ lastResponse, deleteEntry }) => {
    return (
        <div id="response">
            {Object.keys(lastResponse).map((key, id) => (
                <div key={id}>
                    {key.substring(0, 1).toUpperCase() + key.substring(1)}:{" "}
                    {lastResponse[key]}
                </div>
            ))}
            <button
                id="deleteEntry"
                onClick={() => deleteEntry(lastResponse.id)}
            >
                X
            </button>
        </div>
    );
};
