import "../../../../node_modules/terminal.css/dist/terminal.css";
import React, { useState, FormEvent } from 'react';

export default function PostForm() {
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [responseId, setResponseId] = useState<number | null>(null);
  const [error, setError] = useState<string | null>(null);

  async function onSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setIsLoading(true);
    setError(null);
    setResponseId(null);

    const formData = new FormData(event.currentTarget);
    try {
      const response = await fetch('https://localhost:8443/message', {
        method: 'POST',
        body: formData,
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data = await response.json();
      setResponseId(data.id);
    } catch (err) {
      setError(err.message);
    } finally {
      setIsLoading(false);
    }
  }

  return (
    <div className="bg-gray-800 text-white p-12 rounded-lg shadow-lg h-min my-auto terminal-window">
      <h2 className="text-3xl font-bold text-center terminal-window">
        42 Bacheca Di Protesta!
      </h2>
      <form className="flex flex-col gap-6 mt-6" onSubmit={onSubmit}>
        <div>
          <label
            htmlFor="name"
            className="block text-gray-300 font-semibold mb-2 terminal-prompt"
          >
            Name/s of the participants:
          </label>
          <input
            className="w-full h-12 p-4 border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-emerald-500 terminal-input bg-black text-white"
            type="text"
            name="name"
            id="first"
            placeholder="Name/s of the participants (comma separated):"
            required
          />
        </div>
        <div>
          <label
            htmlFor="subject"
            className="block text-gray-300 font-semibold mb-2 terminal-prompt"
          >
            Subject of the request:
          </label>
          <input
            className="w-full h-12 p-4 border border-gray-700 rounded -lg focus:outline-none focus:ring-2 focus:ring-emerald-500 terminal-input bg-black text-white"
            type="text"
            name="subject"
            id="second"
            placeholder="Subject of the request:"
            required
          />
        </div>
        <div>
          <label
            htmlFor="body"
            className="block text-gray-300 font-semibold mb-2 terminal-prompt"
          >
            Body of the Message:
          </label>
          <textarea
            className="w-full h-40 p-4 border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-emerald-500 terminal-input bg-black text-white"
            name="body"
            id="third"
            placeholder="Body of the message"
            rows="10"
            cols="50"
            required
          />
        </div>
        <button
          type="submit"
          className="w-full bg-emerald-600 text-white py-3 rounded-lg hover:bg-emerald-700 transition duration-200 terminal-button"
        >
          {isLoading ? 'Submitting...' : 'Add to the database'}
        </button>
      </form>
      {responseId && (
        <div className="mt-6 text-center text-green-500">
          <p>Form submitted successfully. ID: {responseId}</p>
        </div>
      )}
      {error && (
        <div className="mt-6 text-center text-red-500">
          <p>Error: {error}</p>
        </div>
      )}
    </div>
  );
}