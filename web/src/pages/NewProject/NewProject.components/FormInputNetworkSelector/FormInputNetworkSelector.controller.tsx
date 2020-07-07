import * as PropTypes from "prop-types";
import * as React from "react";

import { SelectView } from "./FormInputNetworkSelector.view";

type SelectProps = {
    options: string[];
    defaultOption: string;
    selectCallback: (e: string) => void;
};

export const FormInputNetworkSelector = ({ options, defaultOption, selectCallback }: SelectProps) => {
    return <SelectView options={options} defaultOption={defaultOption} selectCallback={selectCallback} />;
};

FormInputNetworkSelector.propTypes = {
    options: PropTypes.array,
    defaultOption: PropTypes.string,
    selectCallback: PropTypes.func,
};

FormInputNetworkSelector.defaultProps = {
    options: ["MAINNET", "CARTAGENET"],
    defaultOption: "MAINNET",
    selectCallback: () => {},
};
