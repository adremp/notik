import { UseQueryResult, useQuery } from "@tanstack/react-query";
import { ReactNode } from "react";

export interface WithQueryProps<R> {
  options: Parameters<typeof useQuery<R>>[0];
  loading: ReactNode;
  error: ReactNode;
  children: ReactNode | ((res: UseQueryResult<R>) => ReactNode);
}

const WithQuery = <R,>(props: WithQueryProps<R>) => {
  const res = useQuery(props.options);

  if (res.isError) return props.error;
  if (res.isLoading) return props.loading;

  return props.children instanceof Function
    ? props.children(res)
    : props.children;
};

export default WithQuery;
