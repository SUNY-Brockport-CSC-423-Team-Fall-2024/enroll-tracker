"use client";

import { useRef, useState, useEffect } from "react";
import styles from "./styles.module.css";
import Table from "../table/table";
import { ITableHeader, ITableRow } from "@/app/lib/definitions";

interface SearchInputProps<T extends ITableRow> {
  id: string;
  name: string;
  resultTableHeaders: ITableHeader[];
  queryCallback: (query: string) => Promise<T[]>;
}

const SearchInput = <T extends ITableRow>({
  id,
  name,
  queryCallback,
  resultTableHeaders,
}: SearchInputProps<T>) => {
  const dropdownRef = useRef<HTMLDivElement | null>(null);
  const [isFocused, setIsFocused] = useState<boolean>(false);
  const [results, setResults] = useState<T[]>([]);
  const [query, setQuery] = useState<string>("");

  useEffect(() => {
    const focusIn = () => {
      if (dropdownRef.current) {
        setIsFocused(true);
      }
    };
    const focusOut = () => {
      if (dropdownRef.current) {
        setIsFocused(false);
      }
    };

    document.addEventListener("focusin", focusIn);
    document.addEventListener("focusout", focusOut);
    return () => {
      document.removeEventListener("focusin", focusIn);
      document.removeEventListener("focusout", focusOut);
    };
  }, []);

  useEffect(() => {
    let isLatest = true;
    const runQuery = setTimeout(async () => {
      if (isLatest) {
        console.log("hi");
        const res = await queryCallback(query);

        if (res.length > 0) {
          setResults(res);
        }
      }
    }, 1500);

    return () => {
      clearTimeout(runQuery);
      isLatest = false;
    };
  }, [query, queryCallback]);

  return (
    <div className={styles.search_results_container}>
      <input
        id={id}
        name={name}
        onChange={(e) => setQuery(e.target.value)}
        onFocus={() => setIsFocused(!isFocused)}
      />
      {isFocused && (
        <div ref={dropdownRef} className={styles.search_results_dropdown}>
          <Table headers={resultTableHeaders} rows={results} />
        </div>
      )}
    </div>
  );
};

export default SearchInput;
