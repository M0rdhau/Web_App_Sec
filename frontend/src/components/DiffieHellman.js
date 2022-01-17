import React from "react";
import { useForm } from "../hooks/useForm";
import encryptionService from "../services/encryptionsService";

export const DiffieHellman = ({ encType, updateHistory }) => {

    const [dhValues, handleDHChange, resetDH] = useForm({
        prime: 0,
        primitive: 0,
        userSecret: 0,
    });

    const handleEntryAdd = async () => {
        const response = await encryptionService.addEntry(
            {
                prime: parseInt(dhValues.prime),
                primitive: parseInt(dhValues.primitive),
                userSecret: parseInt(dhValues.userSecret),
            },
            encType.route
        );
        updateHistory(response.data);
    };

    return (
        <div className="encryptionSubmit">
            {encType.name}:<button onClick={resetDH}>Reset Input</button>
            Prime:{" "}
            <input
                className="prime"
                name="prime"
                value={dhValues.prime}
                onChange={handleDHChange}
            />
            Primitive:{" "}
            <input
                className="primitive"
                name="primitive"
                value={dhValues.primitive}
                onChange={handleDHChange}
            />
            UserSecret:{" "}
            <input
                className="userSecret"
                name="userSecret"
                value={dhValues.userSecret}
                onChange={handleDHChange}
            />
            <button onClick={handleEntryAdd}>Add</button>
        </div>
    );
};
