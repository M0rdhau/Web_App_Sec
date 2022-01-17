import React from "react";
import { useForm } from "../hooks/useForm";
import encryptionService from "../services/encryptionsService";

export const RSAEntry = ({ encType, updateHistory }) => {

    const [rsaValues, handleRSAChange, resetRSA] = useForm({
        primep: 0,
        primeq: 0,
    });

    const handleEntryAdd = async () => {
        const response = await encryptionService.addEntry(
            {
                primep: parseInt(rsaValues.primep),
                primeq: parseInt(rsaValues.primeq),
            },
            encType.route
        );
        updateHistory(response.data);
    };

    return (
        <div className="encryptionSubmit">
            {encType.name}:<button onClick={resetRSA}>Reset Input</button>
            Prime P:{" "}
            <input
                className="primep"
                name="primep"
                value={rsaValues.primep}
                onChange={handleRSAChange}
            />
            Prime Q:{" "}
            <input
                className="primeq"
                name="primeq"
                value={rsaValues.primeq}
                onChange={handleRSAChange}
            />
            <button onClick={handleEntryAdd}>Add</button>
        </div>
    );
};
