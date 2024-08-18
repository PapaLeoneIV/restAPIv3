import React from "react";

export default function ButtonList() {
  return (
    <div className="flex flex-col  w-1/12 bg-white">
      <button className="bg-emerald-600 hover:bg-emerald-700 transition duration-200 terminal-button">
        GET LATEST
      </button>
      <button className="bg-emerald-600 hover:bg-emerald-700 transition duration-200 terminal-button">
        GET ALL
      </button>
      <button className="bg-emerald-600 hover:bg-emerald-700 transition duration-200 terminal-button">
        POST
      </button>
      <button className="bg-emerald-600 hover:bg-emerald-700 transition duration-200 terminal-button">
        MODIFY
      </button>
      <button className="bg-emerald-600 hover:bg-emerald-700 transition duration-200 terminal-button">
        DELETE
      </button>
    </div>
  );
}
