import { useState } from "react";

export function ShowSingle() {
    /*TODO display the range of id available*/
    /*TODO inserire bottone per invio richiesta sotto l inserimento dell id*/
    /*TODO fetch the data and display it*/

    (async () => {
      const res = await fetch('http://localhost:8080/message');
      if(!res.ok)
      {
        console.log('Error fetching the data');
        return;
      }
      const data = await res.json();
      console.log(data);
    })();

  return (
    <div className="h-auto w-auto bg-gray-800 text-white p-12 rounded-lg shadow-lg terminal-window">
      <h2 className="text-3xl font-bold text-center terminal-window">
        Insert the Message Id you want to retrieve!
      </h2>
      <form>
        <label>
          <input
            className="w-full h-12 p-4 border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-emerald-500 terminal-input bg-black text-white"
            type="text"
            name="id"
            id="id"
            placeholder="Id of the message"
            required
          />
        </label>
      </form>
      <div className="flex justify-center mt-4">
        <button className="rounded-lg bg-emerald-800 text-white px-4 py-2" >
          Send Request!
        </button>
      </div>
    </div>
  );
}
