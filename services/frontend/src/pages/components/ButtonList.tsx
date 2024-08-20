import React from "react";
import { useState, useEffect } from "react";

  export default function ButtonList({ onFormChange }: { onFormChange: (value: string) => void }) { 
    return (
    <div className="flex flex-col bg-transparent space-y-28 w-1/12">
      <button className="bg-emerald-600 hover:bg-emerald-700 transition  duration-200" onClick={() => onFormChange("single")}>
        GET SINGLE MESSAGE
      </button>
      <button className="bg-emerald-600 hover:bg-emerald-700 transition  duration-200" onClick={() => onFormChange("all")}>
        GET ALL
      </button>
      <button className="bg-emerald-600 hover:bg-emerald-700 transition  duration-200" onClick={() => onFormChange("post")}>
        POST
      </button>
      <button className="bg-emerald-600 hover:bg-emerald-700 transition  duration-200" onClick={() => onFormChange("modify")}>
        MODIFY
      </button>
        <button className="bg-emerald-600 hover:bg-emerald-700 transition  duration-200" onClick={() => onFormChange("delete")}>
        DELETE
      </button>
    </div>
  );
}
