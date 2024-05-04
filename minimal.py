import os
import sys
# from langchain_openai import ChatOpenAI
from langchain_community.document_loaders import TextLoader
from langchain.indexes import VectorstoreIndexCreator
import warnings
warnings.filterwarnings('ignore')

os.environ["OPENAI_API_KEY"] = "...................."

ROOT_DIR = os.path.dirname(os.path.abspath(__file__))

loader = TextLoader(ROOT_DIR + "/kamus.txt")
index = VectorstoreIndexCreator().from_loaders([loader])
# loader.load()

query = sys.argv[1]

print(index.query(query))