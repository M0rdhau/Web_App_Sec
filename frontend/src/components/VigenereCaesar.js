import React, { useState } from "react";
import { useForm } from "../hooks/useForm";
import encryptionService from "../services/encryptionsService";

export const VigenereCaesar = ({ encType, updateHistory }) => {
    const [decrypt, checkDecrypt] = useState(false);

    const [caesarValues, handleCaesarChange, resetCaesar] = useForm({
        input: "",
        shift: "",
    });

    const handleEntryAdd = async () => {
        const response = await encryptionService.addEntry(
            {
                input: caesarValues.input,
                shift:
                    encType.name === "Caesar"
                        ? parseInt(caesarValues.shift)
                        : caesarValues.shift,
                encrypt: !decrypt,
            },
            encType.route
        );
        updateHistory(response.data);
    };

    const handleCheckBoxChange = () => {
        checkDecrypt(!decrypt);
    };

    return (
        <div className="encryptionSubmit">
            {encType.name}:<button onClick={resetCaesar}>Reset Input</button>
            Input:{" "}
            <input
                className="input"
                name="input"
                value={caesarValues.input}
                onChange={handleCaesarChange}
            />
            Key:{" "}
            <input
                className="shift"
                name="shift"
                value={caesarValues.shift}
                onChange={handleCaesarChange}
            />
            Decryption?:{" "}
            <input
                className="encrypt"
                name="encrypt"
                type="checkbox"
                value={caesarValues.encrypt}
                onChange={handleCheckBoxChange}
            />
            <button onClick={handleEntryAdd}>Add</button>
        </div>
    );
};
